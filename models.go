package go_collector

import "context"

// StorageType the type of storage that will be used to accumulate data
type StorageType int

const (
	// Channel storage using channels
	Channel StorageType = iota
	// List storage using container/list
	List
	// GodsList storage using github.com/emirpasic/gods
	GodsList
	// Slice storage using slice
	Slice
)

// FlushExec a function to call when an action needs to be performed
type FlushExec[T comparable] func(events []T) error

type iStorage[T comparable] interface {
	Put(e *eventExtended[T]) bool
	Get() []*eventExtended[T]
}

// Collector data collector
type Collector[T comparable] interface {
	// AddAsync adds an object to the data collector without waiting for flushFunc to execute.
	// Returns ErrSendToClose if the drive has been closed.
	AddAsync(ctx context.Context, d T) error
	// AddSync adds an object to the data collector while waiting for flushFunc to execute.
	// Returns ErrSendToClose if the data collector has been closed.
	// Return context.DeadlineExceeded if data will not have time to be processed in flushFunc, but context time runs out.
	AddSync(ctx context.Context, d T) error
	// Stop stops the data collector. After stopping an AddAsync or AddSync call will return ErrSendToClose
	Stop()
}
