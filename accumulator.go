package go_events_accumulator

import (
	"context"
	"go-events-accumulator/storage"
	"sync"
	"sync/atomic"
	"time"
)

const (
	defaultFlushSize     = 500
	defaultFlushInterval = time.Second * 5
)

// FlushExec - a function to call when an action needs to be performed
type FlushExec[T comparable] func(events []T) error

// NewAccumulator ...
func NewAccumulator[T comparable](
	flushSize uint,
	flushInterval time.Duration,
	flushFunc FlushExec[T],
) (*accum[T], error) {
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

	a := &accum[T]{
		flushSize:     int(size),
		flushInterval: interval,
		flushFunc:     flushFunc,

		chEvents: make(chan *extend[T], size),
	}

	a.wgStop.Add(1)
	go a.startFlusher()

	return a, nil
}

type accum[T comparable] struct {
	flushSize     int
	flushInterval time.Duration
	flushFunc     FlushExec[T]

	chEvents      chan *extend[T]
	wgCountEvents sync.WaitGroup

	isClose atomic.Bool
	wgStop  sync.WaitGroup
}

// AddAsync ...
func (a *accum[T]) AddAsync(ctx context.Context, event T) error {
	if a.isClose.Load() {
		return ErrSendToClose
	}

	selfCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	select {
	case <-selfCtx.Done():
		return context.DeadlineExceeded
	default:

	}

	a.wgCountEvents.Add(1)
	a.chEvents <- &extend[T]{
		e:    event,
		done: atomic.Bool{},
	}

	return nil
}

// AddSync ...
func (a *accum[T]) AddSync(ctx context.Context, event T) error {
	if a.isClose.Load() {
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

	e := &extend[T]{
		fallback: ch,
		e:        event,
	}
	a.wgCountEvents.Add(1)
	a.chEvents <- e

	select {
	case err := <-ch:
		return err
	case <-selfCtx.Done():
		e.done.Store(true)
		return context.DeadlineExceeded
	}
}

// Stop ...
func (a *accum[T]) Stop() {
	if a.isClose.Load() {
		return
	}
	a.isClose.Store(true)

	a.wgCountEvents.Wait()
	close(a.chEvents)

	a.wgStop.Wait()
}

func (a *accum[T]) startFlusher() {
	defer a.wgStop.Done()

	flushTicker := time.NewTicker(a.flushInterval)
	defer flushTicker.Stop()

	events := storage.NewEventStorage[*extend[T]](a.flushSize)

	for {
		select {
		case e, ok := <-a.chEvents:
			if !ok {
				a.chEvents = nil
				a.flush(events.Get())
				return
			}

			a.wgCountEvents.Done()

			// skip finished event (context.DeadlineExceeded)
			if e.done.Load() {
				continue
			}

			if events.Put(e) {
				continue
			}

			a.flush(events.Get())
			flushTicker.Reset(a.flushInterval)

		case <-flushTicker.C:
			a.flush(events.Get())
		}
	}
}

func (a *accum[T]) flush(events []*extend[T]) {
	if len(events) == 0 {
		return
	}

	originalEvents := make([]T, 0, len(events))
	for _, e := range events {
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
