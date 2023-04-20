/*
 * Copyright (c) 2023.
 *
 * Developed thanks to Nikita Terentyev (nar10z). Use it for good, and let your code work without problems!
 */

package storage

import (
	"container/list"
	"sync"
)

func NewEventStorage2[T comparable](size int) *eventStorage2[T] {
	return &eventStorage2[T]{
		size:   size,
		events: list.New(),
		data: sync.Pool{New: func() any {
			data := make([]T, 0, size)
			return data
		}},
	}
}

type eventStorage2[T comparable] struct {
	size   int
	events *list.List
	data   sync.Pool
	mu     sync.Mutex
}

func (s *eventStorage2[T]) Put(e T) bool {
	s.mu.Lock()
	s.events.PushBack(e)
	l := s.events.Len()
	s.mu.Unlock()
	return l < s.size
}

func (s *eventStorage2[T]) Get() []T {
	dataPool := s.data.Get()
	data, _ := dataPool.([]T)

	s.mu.Lock()
	for temp := s.events.Front(); temp != nil; temp = temp.Next() {
		data = append(data, temp.Value.(T))
	}
	s.events.Init()
	s.mu.Unlock()
	s.data.Put(dataPool)

	return data
}
