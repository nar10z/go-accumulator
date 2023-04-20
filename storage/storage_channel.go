/*
 * Copyright (c) 2023.
 *
 * Developed thanks to Nikita Terentyev (nar10z). Use it for good, and let your code work without problems!
 */

package storage

import (
	"sync"
)

func NewStorageChannel[T comparable](size int) *storageChannel[T] {
	return &storageChannel[T]{
		size:   size,
		events: make(chan T, size),
		data: sync.Pool{New: func() any {
			data := make([]T, 0, size)
			return data
		}},
	}
}

type storageChannel[T comparable] struct {
	size   int
	events chan T
	data   sync.Pool
}

func (s *storageChannel[T]) Put(e T) bool {
	s.events <- e
	return len(s.events) < s.size
}

func (s *storageChannel[T]) Get() []T {
	dataPool := s.data.Get()
	data, _ := dataPool.([]T)

	l := len(s.events) // fix chan size
	for i := 0; i < l; i++ {
		data = append(data, <-s.events)
	}
	s.data.Put(dataPool)

	return data
}
