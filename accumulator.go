/*
 * Copyright (c) 2023.
 *
 * License MIT (https://raw.githubusercontent.com/nar10z/go-accumulator/main/LICENSE)
 *
 * Developed thanks to Nikita Terentyev (nar10z). Use it for good, and let your code work without problems!
 */

package go_accumulator

import (
	"context"
	"sync"
	"sync/atomic"
	"time"
)

const (
	defaultFlushSize     = 500
	defaultFlushInterval = time.Second * 5
)

// New creates a new data accumulator
func New[T any](
	flushSize uint,
	flushInterval time.Duration,
	flushFunc FlushExec[T],
) (Accumulator[T], error) {
	size := flushSize
	if size == 0 {
		size = defaultFlushSize
	}

	interval := flushInterval
	if interval == 0 {
		interval = defaultFlushInterval
	}

	if flushFunc == nil {
		return nil, ErrNilFlushFunc
	}

	a := &accumulator[T]{
		flushFunc: flushFunc,
		size:      int(flushSize),
		interval:  flushInterval,

		chEvents: make(chan eventExtended[T], size),
		batchEvents: sync.Pool{
			New: func() any {
				ss := make([]eventExtended[T], 0, size)
				return ss
			},
		},
	}

	a.wgStop.Add(1)
	go a.startFlusher()

	return a, nil
}

type accumulator[T any] struct {
	flushFunc FlushExec[T]

	size     int
	interval time.Duration

	chEvents    chan eventExtended[T]
	batchEvents sync.Pool

	isClose atomic.Bool
	wgStop  sync.WaitGroup
}

func (a *accumulator[T]) AddAsync(ctx context.Context, event T) error {
	if err := a.beforeAddCheck(ctx); err != nil {
		return err
	}

	a.chEvents <- eventExtended[T]{e: event}
	return nil
}

func (a *accumulator[T]) AddSync(ctx context.Context, event T) error {
	if err := a.beforeAddCheck(ctx); err != nil {
		return err
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

func (a *accumulator[T]) beforeAddCheck(ctx context.Context) error {
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

func (a *accumulator[T]) Stop() {
	if a.isClose.Load() {
		return
	}
	a.isClose.Store(true)

	close(a.chEvents)

	a.wgStop.Wait()
}

func (a *accumulator[T]) newBatch() []eventExtended[T] {
	ss, _ := a.batchEvents.Get().([]eventExtended[T])
	return ss
}

func (a *accumulator[T]) clearBatch(s []eventExtended[T]) {
	s = s[:0]
	a.batchEvents.Put(s)
}

func (a *accumulator[T]) startFlusher() {
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

func (a *accumulator[T]) flush(events []eventExtended[T]) {
	if len(events) == 0 {
		return
	}

	originalEvents := make([]T, len(events))
	for i := range events {
		originalEvents[i] = events[i].e
	}

	err := a.flushFunc(originalEvents)
	for _, e := range events {
		if e.fallback == nil {
			continue
		}
		e.fallback <- err
	}
}
