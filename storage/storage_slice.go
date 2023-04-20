/*
 * Copyright (c) 2023.
 *
 * Developed thanks to Nikita Terentyev (nar10z). Use it for good, and let your code work without problems!
 */

package storage

import (
	"sync"
)

func NewEventStorage4[T comparable](size int) *eventStorage4[T] {
	return &eventStorage4[T]{
		size:   size,
		events: make([]T, 0, size),
		data: sync.Pool{New: func() any {
			data := make([]T, 0, size)
			return data
		}},
	}
}

type eventStorage4[T comparable] struct {
	size   int
	events []T
	data   sync.Pool
	mu     sync.Mutex
}

func (s *eventStorage4[T]) Put(e T) bool {
	s.mu.Lock()
	s.events = append(s.events, e)
	l := len(s.events)
	s.mu.Unlock()
	return l < s.size
}

func (s *eventStorage4[T]) Get() []T {
	dataPool := s.data.Get()
	data, _ := dataPool.([]T)

	s.mu.Lock()
	data = s.events[:]
	s.events = s.events[:0]
	s.mu.Unlock()
	s.data.Put(dataPool)

	return data
}
