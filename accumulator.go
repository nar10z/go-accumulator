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

type Accumulator[T comparable] interface {
	AddAsync(ctx context.Context, event T) error
	AddSync(ctx context.Context, event T) error
	Stop()
}

type Opts[T comparable] struct {
	FlushSize     uint
	FlushInterval time.Duration
	FlushFunc     FlushExec[T]
}

func NewAccumulator[T comparable](opts Opts[T]) (Accumulator[T], error) {
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

		chStop:   make(chan struct{}),
		chEvents: make(chan *eventExtend[T], size),

		flushFunc: opts.FlushFunc,
	}

	go a.startFlusher()

	return a, nil
}

type accum[T comparable] struct {
	flushSize     int
	flushInterval time.Duration

	chStop   chan struct{}
	chEvents chan *eventExtend[T]

	isClose   bool
	flushFunc FlushExec[T]
}

func (a *accum[T]) AddAsync(ctx context.Context, event T) error {
	if a.isClose {
		return ErrSendToClose
	}

	select {
	case <-ctx.Done():
		return context.DeadlineExceeded
	default:
		a.chEvents <- &eventExtend[T]{
			e:    event,
			done: atomic.Bool{},
		}
	}

	return nil
}

func (a *accum[T]) AddSync(ctx context.Context, event T) error {
	if a.isClose {
		return ErrSendToClose
	}

	ch := make(chan error)
	e := &eventExtend[T]{
		fallback: ch,
		e:        event,
	}
	defer func() {
		e.done.Store(true)
		close(ch)
	}()

	a.chEvents <- e

	select {
	case err := <-ch:
		return err
	case <-ctx.Done():
		return context.DeadlineExceeded
	}
}

func (a *accum[T]) Stop() {
	a.isClose = true
	close(a.chStop)
	close(a.chEvents)
}

func (a *accum[T]) startFlusher() {
	flushTicker := time.NewTicker(a.flushInterval)
	defer flushTicker.Stop()

	wg := sync.WaitGroup{}
	events := make([]*eventExtend[T], 0, a.flushSize)

	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			select {
			case e := <-a.chEvents:
				// skip finished event (context.DeadlineExceeded)
				if e.done.Load() {
					continue
				}

				events = append(events, e)
				if len(events) < a.flushSize {
					continue
				}

				a.flush(events)
				events = events[:0]
				flushTicker.Reset(a.flushInterval)

			case <-flushTicker.C:
				a.flush(events)
				events = events[:0]

			case <-a.chStop:
				a.flush(events)
				return
			}
		}
	}()

	wg.Wait()
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
		if e.fallback == nil || e.done.Load() {
			continue
		}

		e.done.Store(true)
		e.fallback <- err
	}
}
