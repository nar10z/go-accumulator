/*
 * Copyright (c) 2023.
 *
 * Developed thanks to Nikita Terentyev (nar10z). Use it for good, and let your code work without problems!
 */

package go_collector

import (
	"context"
	"sync"
	"sync/atomic"
	"time"

	"github.com/nar10z/go-collector/storage"
)

const (
	defaultFlushSize     = 500
	defaultFlushInterval = time.Second * 5
)

// New создает новый накопитель
func New[T comparable](
	flushSize uint,
	flushInterval time.Duration,
	flushFunc FlushExec[T],
) (Collector[T], error) {
	return NewWithStorage(flushSize, flushInterval, flushFunc, Channel)
}

// NewWithStorage ...
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

		chEvents: make(chan *eventExtended[T], size),
	}

	a.wgStop.Add(1)
	go a.startFlusher(st)

	return a, nil
}

type collector[T comparable] struct {
	flushSize     int
	flushInterval time.Duration
	flushFunc     FlushExec[T]

	chEvents     chan *eventExtended[T]
	wgAddCounter sync.WaitGroup

	isClose atomic.Bool
	wgStop  sync.WaitGroup
}

// AddAsync ...
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

// AddSync ...
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

// Stop ...
func (c *collector[T]) Stop() {
	if c.isClose.Load() {
		return
	}
	c.isClose.Store(true)

	c.wgAddCounter.Wait()
	close(c.chEvents)

	c.wgStop.Wait()
}

func (c *collector[T]) startFlusher(st StorageType) {
	defer c.wgStop.Done()

	flushTicker := time.NewTicker(c.flushInterval)
	defer flushTicker.Stop()

	var events iStorage[T]
	switch st {
	case Channel:
		events = storage.NewStorageChannel[*eventExtended[T]](c.flushSize)
	case GodsList:
		events = storage.NewStorageSinglyList[*eventExtended[T]](c.flushSize)
	default:
		events = storage.NewStorageList[*eventExtended[T]](c.flushSize)
	}

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
