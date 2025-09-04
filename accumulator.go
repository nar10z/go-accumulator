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
	"fmt"
	"sync/atomic"
	"time"

	"github.com/bytedance/gopkg/lang/syncx"
)

const (
	defaultFlushSize     = 1000
	defaultFlushInterval = time.Millisecond * 250
)

// New creates a new data Accumulator
func New[T any](
	flushSize uint,
	flushInterval time.Duration,
	flushTimeout time.Duration,
	flushFunc FlushExec[T],
) *Accumulator[T] {
	if flushSize == 0 {
		flushSize = defaultFlushSize
	}

	if flushInterval == 0 {
		flushInterval = defaultFlushInterval
	}

	if flushTimeout == 0 {
		flushTimeout = flushInterval
	}

	if flushFunc == nil {
		flushFunc = noop[T]
	}

	a := &Accumulator[T]{
		flushFunc:    flushFunc,
		flushTimeout: flushTimeout,

		chEvents: make(chan eventExtended[T], flushSize),
		batchEvents: syncx.Pool{
			New: func() any {
				return make([]eventExtended[T], 0, flushSize)
			},
			NoGC: true,
		},
		batchOrigEvents: syncx.Pool{
			New: func() any {
				return make([]T, 0, flushSize)
			},
			NoGC: true,
		},

		chStop: make(chan struct{}),
	}

	go a.startFlusher(flushInterval, int(flushSize))

	return a
}

type Accumulator[T any] struct {
	batchEvents     syncx.Pool
	batchOrigEvents syncx.Pool

	flushFunc    FlushExec[T]
	flushTimeout time.Duration

	chEvents chan eventExtended[T]
	chStop   chan struct{}

	isClose atomic.Bool
}

func (a *Accumulator[T]) AddAsync(ctx context.Context, event T) (err error) {
	if a.isClose.Load() {
		return ErrSendToClose
	}

	defer func() {
		// recover from panic caused by writing to a closed channel
		if r := recover(); r != nil {
			err = fmt.Errorf("AddSync, recover: %v", r)
		}
	}()

	select {
	case <-ctx.Done():
		return fmt.Errorf("AddAsync, check on write: %w", ctx.Err())
	default:
		a.chEvents <- eventExtended[T]{e: event}
	}

	return nil
}

func (a *Accumulator[T]) AddSync(ctx context.Context, event T) (err error) {
	// check context before alloc eventExtended
	select {
	case <-ctx.Done():
		return fmt.Errorf("AddSync, check before: %w", ctx.Err())
	default:
	}

	e := eventExtended[T]{
		fallback: make(chan error, 1),
		e:        event,
	}

	if a.isClose.Load() {
		return ErrSendToClose
	}

	defer func() {
		// recover from panic caused by writing to a closed channel
		if r := recover(); r != nil {
			err = fmt.Errorf("AddSync, recover: %v", r)
		}
	}()

	// check context with write to channel
	select {
	case <-ctx.Done():
		return fmt.Errorf("AddSync, check on write: %w", ctx.Err())
	case a.chEvents <- e:
	}

	// check context with wait event result
	select {
	case err = <-e.fallback:
		if err != nil {
			return fmt.Errorf("AddSync, check fallback: %w", err)
		}

		return nil
	case <-ctx.Done():
		return fmt.Errorf("AddSync, check fallback: %w", ctx.Err())
	}
}

func (a *Accumulator[T]) Stop() {
	if !a.isClose.CompareAndSwap(false, true) {
		return
	}

	close(a.chEvents)
	<-a.chStop
}

func (a *Accumulator[T]) IsClosed() bool {
	return a.isClose.Load()
}

func (a *Accumulator[T]) startFlusher(interval time.Duration, size int) {
	ticker := time.NewTicker(interval)
	batch, _ := a.batchEvents.Get().([]eventExtended[T])
	flush := func() {
		a.flush(batch)
		a.batchEvents.Put(batch[:0])
		batch, _ = a.batchEvents.Get().([]eventExtended[T])
	}

loop:
	for {
		select {
		case e, ok := <-a.chEvents:
			if !ok {
				break loop
			}

			batch = append(batch, e)
			if len(batch) < size {
				continue
			}

			flush()
		case <-ticker.C:
			flush()
		}
	}

	ticker.Stop()
	flush()
	a.chStop <- struct{}{}
}

func (a *Accumulator[T]) flush(events []eventExtended[T]) {
	if len(events) == 0 {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), a.flushTimeout)
	defer cancel()

	originalEvents, _ := a.batchOrigEvents.Get().([]T)
	for i := 0; i < len(events); i++ {
		originalEvents = append(originalEvents, events[i].e)
	}

	err := a.flushFunc(ctx, originalEvents)
	for i := 0; i < len(events); i++ {
		if events[i].fallback == nil {
			continue
		}

		select {
		case events[i].fallback <- err:
		default:
		}
	}

	a.batchOrigEvents.Put(originalEvents[:0])
}
