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

// New creates a new data accumulator with default storage (Channel)
func New[T comparable](
	flushSize uint,
	flushInterval time.Duration,
	flushFunc FlushExec[T],
) (Accumulator[T], error) {
	return NewWithStorage(flushSize, flushInterval, flushFunc, StdList)
}

// NewWithStorage creates a new data accumulator with the specified storage
func NewWithStorage[T comparable](
	flushSize uint,
	flushInterval time.Duration,
	flushFunc FlushExec[T],
	st StorageType,
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
	}

	switch st {
	case Slice:
		a.storage = storage.NewStorageSlice[*eventExtended[T]](int(size))
	case List:
		a.storage = storage.NewStorageSinglyList[*eventExtended[T]](int(size))
	case StdList:
		a.storage = storage.NewStorageList[*eventExtended[T]](int(size))
	case Channel:
		a.storage = storage.NewStorageChannel[*eventExtended[T]](int(size))
	default:
		return nil, ErrNotSetStorageType
	}

	a.wgStop.Add(1)
	go a.startFlusher(interval)

	return a, nil
}

type accumulator[T comparable] struct {
	flushFunc FlushExec[T]

	chEvents chan *eventExtended[T]
	storage  iStorage[T]

	isClose atomic.Bool
	wgStop  sync.WaitGroup
}

func (a *accumulator[T]) AddAsync(ctx context.Context, event T) error {
	if a.isClose.Load() {
		return ErrSendToClose
	}

	select {
	case <-ctx.Done():
		return context.DeadlineExceeded
	default:

	}

	a.chEvents <- &eventExtended[T]{e: event}

	return nil
}

func (a *accumulator[T]) AddSync(ctx context.Context, event T) error {
	if a.isClose.Load() {
		return ErrSendToClose
	}

	ch := make(chan error)
	defer close(ch)

	select {
	case <-ctx.Done():
		return context.DeadlineExceeded
	default:

	}

	e := &eventExtended[T]{
		fallback: ch,
		e:        event,
	}
	a.chEvents <- e

	select {
	case err := <-ch:
		return err
	case <-ctx.Done():
		e.done.Store(true)
		return context.DeadlineExceeded
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

	flushTicker := time.NewTicker(flushInterval)
	defer flushTicker.Stop()

	for {
		select {
		case e, ok := <-a.chEvents:
			if !ok {
				a.chEvents = nil
				a.flush(a.storage.Get())
				return
			}

			// skip finished event (context.DeadlineExceeded)
			if e.done.Load() {
				continue
			}

			if a.storage.Put(e) {
				continue
			}

			a.flush(a.storage.Get())
			flushTicker.Reset(flushInterval)

		case <-flushTicker.C:
			a.flush(a.storage.Get())
		}
	}
}

func (a *accumulator[T]) flush(events []*eventExtended[T]) {
	if len(events) == 0 {
		return
	}

	originalEvents := make([]T, 0, len(events))
	for _, e := range events {
		if e.done.Load() {
			continue
		}
		originalEvents = append(originalEvents, e.e)
	}

	err := a.flushFunc(originalEvents)
	for _, e := range events {
		isDone := e.done.Load()
		e.done.Store(true)

		if isDone || e.fallback == nil {
			continue
		}
		e.fallback <- err
	}
}
