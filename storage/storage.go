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
)

// New creates a new storage
func New[T any](maxSize int) *Storage[T] {
	s := &Storage[T]{
		maxSize: maxSize,
	}
	s.buildEvents()

	return s
}

// Storage ...
type Storage[T any] struct {
	maxSize int
	events  []T
	data    sync.Pool
}

func (s *Storage[T]) buildEvents() {
	s.events, _ = s.data.Get().([]T)
}

// Put add a new data
func (s *Storage[T]) Put(e T) bool {
	s.events = append(s.events, e)
	return len(s.events) < s.maxSize
}

// Len returns size on storage
func (s *Storage[T]) Len() int {
	return len(s.events)
}

// Get returns all data in storage
func (s *Storage[T]) Get() []T {
	return s.events
}

// Clear reset storage
func (s *Storage[T]) Clear() {
	s.events = s.events[:0]
	s.data.Put(s.events)
	s.buildEvents()
}
