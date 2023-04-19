package go_events_accumulator

import (
	"sync"
)

func newEventStorage[T comparable](size int) *eventStorage[T] {
	return &eventStorage[T]{
		size:   size,
		events: make(chan *eventExtend[T], size),
		mu:     sync.Mutex{},
	}
}

type eventStorage[T comparable] struct {
	size   int
	events chan *eventExtend[T]
	mu     sync.Mutex
}

func (s *eventStorage[T]) put(e *eventExtend[T]) bool {
	s.events <- e
	return len(s.events) < s.size
}

func (s *eventStorage[T]) get() []*eventExtend[T] {
	data := make([]*eventExtend[T], 0, s.size)
	l := len(s.events)
	for i := 0; i < l; i++ {
		data = append(data, <-s.events)
	}

	return data
}
