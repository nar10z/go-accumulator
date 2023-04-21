//go:build go1.20

/*
 * Copyright (c) 2023.
 *
 * License MIT (https://raw.githubusercontent.com/nar10z/go-collector/main/LICENSE)
 *
 * Developed thanks to Nikita Terentyev (nar10z). Use it for good, and let your code work without problems!
 */

package go_collector

import "github.com/nar10z/go-collector/storage"

// StorageType the type of storage that will be used to accumulate data
type StorageType int

const (
	// Channel storage using channels
	Channel StorageType = iota
	// List storage using github.com/emirpasic/gods
	List
	// Slice storage using slice
	Slice
	// StdList storage using container/list
	StdList
)

type iStorage[T comparable] interface {
	Put(e *eventExtended[T]) bool
	Get() []*eventExtended[T]
}

func buildStorage[T comparable](st StorageType, flushSize int) iStorage[T] {
	switch st {
	case Channel:
		return storage.NewStorageChannel[*eventExtended[T]](flushSize)
	case List:
		return storage.NewStorageSinglyList[*eventExtended[T]](flushSize)
	case Slice:
		return storage.NewStorageSlice[*eventExtended[T]](flushSize)
	case StdList:
		return storage.NewStorageList[*eventExtended[T]](flushSize)
	default:
		return storage.NewStorageChannel[*eventExtended[T]](flushSize)
	}
}
