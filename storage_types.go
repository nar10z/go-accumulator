/*
 * Copyright (c) 2023.
 *
 * License MIT (https://raw.githubusercontent.com/nar10z/go-accumulator/main/LICENSE)
 *
 * Developed thanks to Nikita Terentyev (nar10z). Use it for good, and let your code work without problems!
 */

package go_accumulator

// StorageType the type of storage that will be used to accumulate data
type StorageType int

const (

	// List storage using github.com/emirpasic/gods
	List StorageType = iota
	// Slice storage using slice
	Slice
	// StdList storage using container/list
	StdList
)

type iStorage[T comparable] interface {
	Put(e *eventExtended[T]) bool
	Len() int
	Iterate(func(ee *eventExtended[T]))
	Clear()
}
