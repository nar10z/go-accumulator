/*
 * Copyright (c) 2023.
 *
 * License MIT (https://raw.githubusercontent.com/nar10z/go-collector/main/LICENSE)
 *
 * Developed thanks to Nikita Terentyev (nar10z). Use it for good, and let your code work without problems!
 */

package go_collector

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

// New creates a new data collector with default storage (Channel)
func New[T comparable](
	flushSize uint,
	flushInterval time.Duration,
	flushFunc FlushExec[T],
) (Collector[T], error) {
	return NewWithStorage(flushSize, flushInterval, flushFunc, Channel)
}

// NewWithStorage creates a new data collector with the specified storage
func NewWithStorage[T comparable](
	flushSize uint,
	flushInterval time.Duration,
	flushFunc FlushExec[T],
	st StorageType,
) (Collector[T], error) {
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

	a := &collector[T]{
		flushSize:     int(size),
		flushInterval: interval,
		flushFunc:     flushFunc,
		storageType:   st,

		chEvents: make(chan *eventExtended[T], size),
	}

	a.wgStop.Add(1)
	go a.startFlusher()

	return a, nil
}

type collector[T comparable] struct {
	flushSize     int
	flushInterval time.Duration
	flushFunc     FlushExec[T]
	storageType   StorageType

	chEvents     chan *eventExtended[T]
	wgAddCounter sync.WaitGroup

	isClose atomic.Bool
	wgStop  sync.WaitGroup
}

func (c *collector[T]) AddAsync(ctx context.Context, event T) error {
	if c.isClose.Load() {
		return ErrSendToClose
	}

	selfCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	select {
	case <-selfCtx.Done():
		return context.DeadlineExceeded
	default:

	}

	c.wgAddCounter.Add(1)
	c.chEvents <- &eventExtended[T]{
		e:    event,
		done: atomic.Bool{},
	}

	return nil
}

func (c *collector[T]) AddSync(ctx context.Context, event T) error {
	if c.isClose.Load() {
		return ErrSendToClose
	}

	selfCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	ch := make(chan error)
	defer close(ch)

	select {
	case <-selfCtx.Done():
		return context.DeadlineExceeded
	default:

	}

	e := &eventExtended[T]{
		fallback: ch,
		e:        event,
	}
	c.wgAddCounter.Add(1)
	c.chEvents <- e

	select {
	case err := <-ch:
		return err
	case <-selfCtx.Done():
		e.done.Store(true)
		return context.DeadlineExceeded
	}
}

func (c *collector[T]) Stop() {
	if c.isClose.Load() {
		return
	}
	c.isClose.Store(true)

	c.wgAddCounter.Wait()
	close(c.chEvents)

	c.wgStop.Wait()
}

func (c *collector[T]) startFlusher() {
	defer c.wgStop.Done()

	flushTicker := time.NewTicker(c.flushInterval)
	defer flushTicker.Stop()

	events := buildStorage[T](c.storageType, c.flushSize)

	for {
		select {
		case e, ok := <-c.chEvents:
			if !ok {
				c.chEvents = nil
				c.flush(events.Get())
				return
			}

			c.wgAddCounter.Done()

			// skip finished event (context.DeadlineExceeded)
			if e.done.Load() {
				continue
			}

			if events.Put(e) {
				continue
			}

			c.flush(events.Get())
			flushTicker.Reset(c.flushInterval)

		case <-flushTicker.C:
			c.flush(events.Get())
		}
	}
}

func (c *collector[T]) flush(events []*eventExtended[T]) {
	if len(events) == 0 {
		return
	}

	originalEvents := make([]T, 0, len(events))
	for _, e := range events {
		originalEvents = append(originalEvents, e.e)
	}

	err := c.flushFunc(originalEvents)
	for _, e := range events {
		isDone := e.done.Load()
		e.done.Store(true)

		if isDone || e.fallback == nil {
			continue
		}
		e.fallback <- err
	}
}
