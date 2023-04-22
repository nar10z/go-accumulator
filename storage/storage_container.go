/*
 * Copyright (c) 2023.
 *
 * License MIT (https://raw.githubusercontent.com/nar10z/go-accumulator/main/LICENSE)
 *
 * Developed thanks to Nikita Terentyev (nar10z). Use it for good, and let your code work without problems!
 */

package storage

import (
	"container/list"
	"sync"
)

// NewStorageList creates a new storage that uses the container/list
func NewStorageList[T comparable](size int) *storageList[T] {
	return &storageList[T]{
		size:   size,
		events: list.New(),
		data: sync.Pool{New: func() any {
			data := make([]T, 0, size)
			return data
		}},
	}
}

type storageList[T comparable] struct {
	size   int
	events *list.List
	data   sync.Pool
	mu     sync.Mutex
}

func (s *storageList[T]) Put(e T) bool {
	s.mu.Lock()
	s.events.PushBack(e)
	l := s.events.Len()
	s.mu.Unlock()
	return l < s.size
}

func (s *storageList[T]) Len() int {
	return s.events.Len()
}

func (s *storageList[T]) Iterate(f func(ee T)) {
	s.mu.Lock()
	for temp := s.events.Front(); temp != nil; temp = temp.Next() {
		f(temp.Value.(T))
	}
	s.mu.Unlock()
}

func (s *storageList[T]) Clear() {
	s.events.Init()
}
