package storage

import (
	"sync"
)

func NewEventStorage[T comparable](size int) *eventStorage[T] {
	return &eventStorage[T]{
		size:   size,
		events: make(chan T, size),
		mu:     sync.Mutex{},
	}
}

type eventStorage[T comparable] struct {
	size   int
	events chan T
	mu     sync.Mutex
}

func (s *eventStorage[T]) Put(e T) bool {
	s.events <- e
	return len(s.events) < s.size
}

func (s *eventStorage[T]) Get() []T {
	l := len(s.events) // fix chan size
	data := make([]T, 0, l)
	for i := 0; i < l; i++ {
		data = append(data, <-s.events)
	}

	return data
}
