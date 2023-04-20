/*
 * Copyright (c) 2023.
 *
 * Developed thanks to Nikita Terentyev (nar10z). Use it for good, and let your code work without problems!
 */

package storage

import (
	"sync"

	sll "github.com/emirpasic/gods/lists/singlylinkedlist"
)

func NewEventStorage3[T comparable](size int) *eventStorage3[T] {
	return &eventStorage3[T]{
		size:   size,
		events: sll.New(),
		data: sync.Pool{New: func() any {
			data := make([]T, 0, size)
			return data
		}},
	}
}

type eventStorage3[T comparable] struct {
	size   int
	events *sll.List
	data   sync.Pool
	mu     sync.Mutex
}

func (s *eventStorage3[T]) Put(e T) bool {
	s.mu.Lock()
	s.events.Add(e)
	l := s.events.Size()
	s.mu.Unlock()
	return l < s.size
}

func (s *eventStorage3[T]) Get() []T {
	dataPool := s.data.Get()
	data, _ := dataPool.([]T)

	s.mu.Lock()
	s.events.Each(func(_ int, temp any) {
		data = append(data, temp.(T))
	})
	s.events.Clear()
	s.mu.Unlock()
	s.data.Put(dataPool)

	return data
}
