/*
 * Copyright (c) 2023.
 *
 * License MIT (https://raw.githubusercontent.com/nar10z/go-accumulator/main/LICENSE)
 *
 * Developed thanks to Nikita Terentyev (nar10z). Use it for good, and let your code work without problems!
 */

package storage

import (
	"sync"
	"sync/atomic"
)

// NewStorageSlice creates a new storage that uses the slice
func NewStorageSlice[T comparable](maxSize int) *storageSlice[T] {
	s := &storageSlice[T]{
		maxSize: int32(maxSize),
		events:  make([]T, 0, maxSize),
	}
	s.buildEvents()

	return s
}

type storageSlice[T comparable] struct {
	maxSize int32
	events  []T
	size    atomic.Int32
	data    sync.Pool
	mu      sync.Mutex
}

func (s *storageSlice[T]) buildEvents() {
	s.events, _ = s.data.Get().([]T)
}

func (s *storageSlice[T]) Put(e T) bool {
	s.mu.Lock()
	s.events = append(s.events, e)
	s.mu.Unlock()
	return s.size.Add(1) < s.maxSize
}

func (s *storageSlice[T]) Len() int {
	return int(s.size.Load())
}

func (s *storageSlice[T]) Iterate(f func(ee T)) {
	s.mu.Lock()
	for _, event := range s.events {
		f(event)
	}
	s.mu.Unlock()
}

func (s *storageSlice[T]) Clear() {
	s.mu.Lock()
	s.events = make([]T, 0, s.maxSize)
	s.events = s.events[:0]
	s.size.Store(0)
	s.data.Put(s.events)
	s.buildEvents()
	s.mu.Unlock()
}
