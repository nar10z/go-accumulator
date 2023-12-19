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
	"unsafe"
)

const (
	defaultFlushSize     = 1000
	defaultFlushInterval = time.Millisecond * 250
)

// FlushExec a function to call when an action needs to be performed
type FlushExec[T any] func(events []T) error

func noop[T any](_ []T) error {
	return nil
}

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

		chEvents: make(chan eventExtended, flushSize),
		batchEvents: sync.Pool{
			New: func() any {
				ss := make([]eventExtended, 0, flushSize)
				return ss
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

	chEvents    chan eventExtended
	batchEvents sync.Pool

	isClose atomic.Bool
	wgStop  sync.WaitGroup
}

func (a *Accumulator[T]) AddAsync(ctx context.Context, event T) error {
	if err := a.beforeAddCheck(ctx); err != nil {
		return err
	}

	a.chEvents <- eventExtended{e: uintptr(unsafe.Pointer(&event))}
	return nil
}

func (a *Accumulator[T]) AddSync(ctx context.Context, event T) error {
	if err := a.beforeAddCheck(ctx); err != nil {
		return err
	}

	e := eventExtended{
		fallback: make(chan error),
		e:        uintptr(unsafe.Pointer(&event)),
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

func (a *Accumulator[T]) beforeAddCheck(ctx context.Context) error {
	if a.isClose.Load() {
		return ErrSendToClose
	}

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		return nil
	}
}

func (a *Accumulator[T]) Stop() {
	if a.isClose.Load() {
		return
	}
	a.isClose.Store(true)

	close(a.chEvents)

	a.wgStop.Wait()
}

func (a *Accumulator[T]) newBatch() []eventExtended {
	ss, _ := a.batchEvents.Get().([]eventExtended)
	return ss
}

func (a *Accumulator[T]) clearBatch(s []eventExtended) {
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

func (a *Accumulator[T]) flush(events []eventExtended) {
	if len(events) == 0 {
		return
	}

	originalEvents := make([]T, len(events))
	for i := range events {
		originalEvents[i] = *(*T)(unsafe.Pointer(events[i].e))
	}

	err := a.flushFunc(originalEvents)
	for _, e := range events {
		if e.fallback == nil {
			continue
		}

		e.fallback <- err
	}
}
