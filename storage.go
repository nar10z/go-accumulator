package go_events_accumulator

import "sync"

type eventStorage[T comparable] struct {
	size   int
	events []*eventExtend[T]
	mu     sync.Mutex
}

func newEventStorage[T comparable](size int) *eventStorage[T] {
	return &eventStorage[T]{
		size:   size,
		events: make([]*eventExtend[T], 0, size),
		mu:     sync.Mutex{},
	}
}

func (s *eventStorage[T]) put(e *eventExtend[T]) int {
	s.mu.Lock()
	s.events = append(s.events, e)
	l := len(s.events)
	s.mu.Unlock()
	return l
}

func (s *eventStorage[T]) get() []*eventExtend[T] {
	size := s.size

	s.mu.Lock()
	if len(s.events) < size {
		size = len(s.events)
	}

	data := s.events[:size]
	s.events = s.events[size:]
	s.mu.Unlock()

	return data
}
