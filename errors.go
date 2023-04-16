package go_events_accumulator

import "errors"

var (
	// ErrNilFlushFunc ...
	ErrNilFlushFunc = errors.New("nil flush func")
	// ErrSendToClose ...
	ErrSendToClose = errors.New("send to close accumulator")
)
