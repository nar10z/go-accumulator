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

	sll "github.com/emirpasic/gods/lists/singlylinkedlist"
)

// NewStorageSinglyList creates a new storage that uses the github.com/emirpasic/gods/lists/singlylinkedlist
func NewStorageSinglyList[T comparable](size int) *storageSinglyList[T] {
	return &storageSinglyList[T]{
		size:   size,
		events: sll.New(),
	}
}

type storageSinglyList[T comparable] struct {
	size   int
	events *sll.List
	mu     sync.Mutex
}

func (s *storageSinglyList[T]) Put(e T) bool {
	s.mu.Lock()
	s.events.Add(e)
	l := s.events.Size()
	s.mu.Unlock()
	return l < s.size
}

func (s *storageSinglyList[T]) Len() int {
	s.mu.Lock()
	l := s.events.Size()
	s.mu.Unlock()
	return l
}

func (s *storageSinglyList[T]) Iterate(f func(ee T)) {
	s.mu.Lock()
	s.events.Each(func(_ int, temp any) {
		f(temp.(T))
	})
	s.mu.Unlock()
}

func (s *storageSinglyList[T]) Clear() {
	s.mu.Lock()
	s.events.Clear()
	s.mu.Unlock()
}
