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
	"runtime"
	"sync"
)

// NewStorageList creates a new storage that uses the container/list
func NewStorageList[T comparable](size int) *storageList[T] {
	return &storageList[T]{
		size:   size,
		sema:   make(chan struct{}, runtime.GOMAXPROCS(0)),
		events: list.New(),
	}
}

type storageList[T comparable] struct {
	size   int
	sema   chan struct{}
	events *list.List
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
	s.mu.Lock()
	l := s.events.Len()
	s.mu.Unlock()
	return l
}

func (s *storageList[T]) Iterate(f func(ee T)) {
	var wg sync.WaitGroup

	s.mu.Lock()
	for temp := s.events.Front(); temp != nil; temp = temp.Next() {
		wg.Add(1)
		s.sema <- struct{}{}
		event := temp.Value.(T)
		go func() {
			f(event)
			<-s.sema
			wg.Done()
		}()
	}
	wg.Wait()
	s.mu.Unlock()
}

func (s *storageList[T]) Clear() {
	s.events.Init()
}
