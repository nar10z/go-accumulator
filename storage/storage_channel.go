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

// NewStorageChannel creates a new storage that uses the channel
func NewStorageChannel[T comparable](size int) *storageChannel[T] {
	return &storageChannel[T]{
		maxSize: int32(size),
		events:  make(chan T, size),
		data: sync.Pool{New: func() any {
			return make([]T, 0, size)
		}},
	}
}

type storageChannel[T comparable] struct {
	maxSize int32
	events  chan T
	size    atomic.Int32
	data    sync.Pool
}

func (s *storageChannel[T]) Put(e T) bool {
	s.events <- e
	return s.size.Add(1) < s.maxSize
}

func (s *storageChannel[T]) Get() []T {
	dataPool := s.data.Get()
	data, _ := dataPool.([]T)

	l := int(s.size.Swap(0)) // fix chan maxSize
	for i := 0; i < l; i++ {
		data = append(data, <-s.events)
	}
	s.data.Put(dataPool)

	return data
}
