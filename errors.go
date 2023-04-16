package go_events_accumulator

import "errors"

var (
	ErrNilFlushFunc = errors.New("nil flush func")
	ErrSendToClose  = errors.New("send to close accumulator")
)
