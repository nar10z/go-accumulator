package go_events_accumulator

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

// FlushExec - a function to call when an action needs to be performed
type FlushExec[T comparable] func(events []T) error

// Opts ...
type Opts[T comparable] struct {
	FlushSize     uint
	FlushInterval time.Duration
	FlushFunc     FlushExec[T]
}

// NewAccumulator ...
func NewAccumulator[T comparable](opts Opts[T]) (*accum[T], error) {
	size := opts.FlushSize
	if size == 0 {
		size = defaultFlushSize
	}

	interval := opts.FlushInterval
	if interval == 0 {
		interval = defaultFlushInterval
	}

	if opts.FlushFunc == nil {
		return nil, ErrNilFlushFunc
	}

	a := &accum[T]{
		flushSize:     int(size),
		flushInterval: interval,
		flushFunc:     opts.FlushFunc,

		chStop:   make(chan struct{}),
		chEvents: make(chan *eventExtend[T], size),
		events:   newEventStorage[T](int(size)),
	}

	a.wgStop.Add(1)
	go a.startFlusher()

	return a, nil
}

type accum[T comparable] struct {
	flushSize     int
	flushInterval time.Duration
	flushFunc     FlushExec[T]

	chStop        chan struct{}
	chEvents      chan *eventExtend[T]
	counterEvents atomic.Int32
	events        *eventStorage[T]

	isClose atomic.Bool
	wgStop  sync.WaitGroup
}

// AddAsync ...
func (a *accum[T]) AddAsync(ctx context.Context, event T) error {
	if a.isClose.Load() {
		return ErrSendToClose
	}

	select {
	case <-ctx.Done():
		return context.DeadlineExceeded
	default:
		a.counterEvents.Add(1)
		a.chEvents <- &eventExtend[T]{
			e:    event,
			done: atomic.Bool{},
		}
	}

	return nil
}

// AddSync ...
func (a *accum[T]) AddSync(ctx context.Context, event T) error {
	ch := make(chan error)
	defer close(ch)

	e := &eventExtend[T]{
		fallback: ch,
		e:        event,
	}

	if a.isClose.Load() {
		return ErrSendToClose
	}
	a.counterEvents.Add(1)
	a.chEvents <- e

	select {
	case err := <-ch:
		return err
	case <-ctx.Done():
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

	for a.counterEvents.Load() > 0 {
		time.Sleep(time.Microsecond)
	}
	close(a.chEvents)

	a.chStop <- struct{}{}
	a.wgStop.Wait()
	close(a.chStop)
}

func (a *accum[T]) startFlusher() {
	defer a.wgStop.Done()

	flushTicker := time.NewTicker(a.flushInterval)
	defer flushTicker.Stop()

	for {
		select {
		case e, ok := <-a.chEvents:
			if !ok {
				a.chEvents = nil
				continue
			}

			a.counterEvents.Add(-1)

			// skip finished event (context.DeadlineExceeded)
			if e.done.Load() {
				continue
			}

			if a.events.put(e) < a.flushSize {
				continue
			}

			a.flush(a.events.get())
			flushTicker.Reset(a.flushInterval)

		case <-flushTicker.C:
			a.flush(a.events.get())

		case <-a.chStop:
			a.flush(a.events.get())

			return
		}
	}
}

func (a *accum[T]) flush(events []*eventExtend[T]) {
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
