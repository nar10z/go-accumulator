package go_events_accumulator

import "sync/atomic"

type eventExtend[T comparable] struct {
	// return error of flush operation
	fallback chan<- error
	// original data
	e T

	done atomic.Bool
}
