/*
 * Copyright (c) 2023.
 *
 * License MIT (https://raw.githubusercontent.com/nar10z/go-accumulator/main/LICENSE)
 *
 * Developed thanks to Nikita Terentyev (nar10z). Use it for good, and let your code work without problems!
 */

package goaccum

import (
	"context"
	"sync"
	"sync/atomic"
	"time"
)

const (
	defaultFlushSize     = 1000
	defaultFlushInterval = time.Millisecond * 250
)

// New creates a new data Accumulator
func New[T any](
	flushSize uint,
	flushInterval time.Duration,
	flushFunc FlushExec[T],
) *Accumulator[T] {
	if flushSize == 0 {
		flushSize = defaultFlushSize
	}

	if flushInterval == 0 {
		flushInterval = defaultFlushInterval
	}

	if flushFunc == nil {
		flushFunc = noop[T]
	}

	a := &Accumulator[T]{
		flushFunc: flushFunc,
		size:      int(flushSize),
		interval:  flushInterval,

		chEvents: make(chan eventExtended[T], flushSize),
		batchEvents: sync.Pool{
			New: func() any {
				return make([]eventExtended[T], 0, flushSize)
			},
		},
		batchOrigEvents: sync.Pool{
			New: func() any {
				return make([]T, 0, flushSize)
			},
		},
	}

	a.wgStop.Add(1)
	go a.startFlusher()

	return a
}

type Accumulator[T any] struct {
	flushFunc FlushExec[T]

	size     int
	interval time.Duration

	chEvents        chan eventExtended[T]
	batchEvents     sync.Pool
	batchOrigEvents sync.Pool

	isClose atomic.Bool
	wgStop  sync.WaitGroup
}

func (a *Accumulator[T]) AddAsync(ctx context.Context, event T) error {
	if a.isClose.Load() {
		return ErrSendToClose
	}

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		a.chEvents <- eventExtended[T]{e: event}
	}

	return nil
}

func (a *Accumulator[T]) AddSync(ctx context.Context, event T) error {
	if a.isClose.Load() {
		return ErrSendToClose
	}

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	e := eventExtended[T]{
		fallback: make(chan error),
		e:        event,
	}
	a.chEvents <- e

	select {
	case err := <-e.fallback:
		return err
	case <-ctx.Done():
		e.fallback = nil
		return ctx.Err()
	}
}

func (a *Accumulator[T]) Stop() {
	if !a.isClose.CompareAndSwap(false, true) {
		return
	}

	close(a.chEvents)
	a.wgStop.Wait()
}

func (a *Accumulator[T]) newBatch() []eventExtended[T] {
	ss, _ := a.batchEvents.Get().([]eventExtended[T])
	return ss
}

func (a *Accumulator[T]) clearBatch(s []eventExtended[T]) {
	a.batchEvents.Put(s[:0])
}

func (a *Accumulator[T]) startFlusher() {
	defer a.wgStop.Done()

	ticker := time.NewTicker(a.interval)
	defer ticker.Stop()

	batch := a.newBatch()

	flush := func() {
		a.flush(batch)
		a.clearBatch(batch)
		batch = a.newBatch()
	}

	for {
		select {
		case e, ok := <-a.chEvents:
			if !ok {
				a.chEvents = nil
				flush()
				return
			}

			batch = append(batch, e)
			if len(batch) < a.size {
				continue
			}

			flush()

			ticker.Reset(a.interval)
		case <-ticker.C:
			flush()
		}
	}
}

func (a *Accumulator[T]) flush(events []eventExtended[T]) {
	if len(events) == 0 {
		return
	}

	originalEvents, _ := a.batchOrigEvents.Get().([]T)
	for i := range events {
		originalEvents = append(originalEvents, events[i].e)
	}

	err := a.flushFunc(originalEvents)
	for _, e := range events {
		if e.fallback == nil {
			continue
		}

		e.fallback <- err
	}

	a.batchOrigEvents.Put(originalEvents[:0])
}
