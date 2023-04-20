package go_collector

import "context"

type StorageType int

const (
	Channel StorageType = iota
	List
	GodsList
)

// FlushExec - a function to call when an action needs to be performed
type FlushExec[T comparable] func(events []T) error

type iStorage[T comparable] interface {
	Put(e *eventExtended[T]) bool
	Get() []*eventExtended[T]
}

// Collector ...
type Collector[T comparable] interface {
	AddAsync(ctx context.Context, event T) error
	AddSync(ctx context.Context, event T) error
	Stop()
}
