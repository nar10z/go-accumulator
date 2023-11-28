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

	"github.com/nar10z/go-accumulator/storage"
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
		chEvents:  make(chan *eventExtended[T], size),
		storage:   storage.New[*eventExtended[T]](int(size)),
	}

	a.wgStop.Add(1)
	go a.startFlusher(interval)

	return a, nil
}

type accumulator[T any] struct {
	flushFunc FlushExec[T]

	chEvents chan *eventExtended[T]
	storage  *storage.Storage[*eventExtended[T]]

	isClose atomic.Bool
	wgStop  sync.WaitGroup
}

func (a *accumulator[T]) AddAsync(ctx context.Context, event T) error {
	if err := a.beforeAddCheck(ctx); err != nil {
		return err
	}

	a.chEvents <- &eventExtended[T]{e: event}
	return nil
}

func (a *accumulator[T]) AddSync(ctx context.Context, event T) error {
	if err := a.beforeAddCheck(ctx); err != nil {
		return err
	}

	e := &eventExtended[T]{
		fallback: make(chan error),
		e:        event,
	}
	a.chEvents <- e

	select {
	case err := <-e.fallback:
		return err
	case <-ctx.Done():
		e.fallback = nil
		return context.DeadlineExceeded
	}
}

func (a *accumulator[T]) beforeAddCheck(ctx context.Context) error {
	if a.isClose.Load() {
		return ErrSendToClose
	}

	select {
	case <-ctx.Done():
		return context.DeadlineExceeded
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

func (a *accumulator[T]) startFlusher(flushInterval time.Duration) {
	defer a.wgStop.Done()

	ticker := time.NewTicker(flushInterval)
	defer ticker.Stop()

	for {
		select {
		case e, ok := <-a.chEvents:
			if !ok {
				a.chEvents = nil
				a.flush()
				return
			}

			if a.storage.Put(e) {
				continue
			}

			a.flush()
			ticker.Reset(flushInterval)
		case <-ticker.C:
			a.flush()
		}
	}
}

func (a *accumulator[T]) flush() {
	l := a.storage.Len()
	if l == 0 {
		return
	}

	events := a.storage.Get()
	a.storage.Clear()

	originalEvents := make([]T, 0, l)
	for _, e := range events {
		originalEvents = append(originalEvents, e.e)
	}

	err := a.flushFunc(originalEvents)
	for _, e := range events {
		if e.fallback == nil {
			continue
		}
		e.fallback <- err
	}
}
