/*
 * Copyright (c) 2023.
 *
 * License MIT (https://raw.githubusercontent.com/nar10z/go-collector/main/LICENSE)
 *
 * Developed thanks to Nikita Terentyev (nar10z). Use it for good, and let your code work without problems!
 */

package storage

import (
	"sync"
)

// NewStorageSlice creates a new storage that uses the slice
func NewStorageSlice[T comparable](size int) *storageSlice[T] {
	return &storageSlice[T]{
		size:   size,
		events: make([]T, 0, size),
		data: sync.Pool{New: func() any {
			data := make([]T, 0, size)
			return data
		}},
	}
}

type storageSlice[T comparable] struct {
	size   int
	events []T
	data   sync.Pool
	mu     sync.Mutex
}

func (s *storageSlice[T]) Put(e T) bool {
	s.mu.Lock()
	s.events = append(s.events, e)
	l := len(s.events)
	s.mu.Unlock()
	return l < s.size
}

func (s *storageSlice[T]) Get() []T {
	dataPool := s.data.Get()
	data, _ := dataPool.([]T)

	s.mu.Lock()
	data = s.events[:]
	s.events = s.events[:0]
	s.mu.Unlock()
	s.data.Put(dataPool)

	return data
}
