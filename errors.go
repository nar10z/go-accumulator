/*
 * Copyright (c) 2023.
 *
 * Developed thanks to Nikita Terentyev (nar10z). Use it for good, and let your code work without problems!
 */

package go_collector

import "errors"

var (
	// ErrNilFlushFunc ...
	ErrNilFlushFunc = errors.New("nil flush func")
	// ErrSendToClose ...
	ErrSendToClose = errors.New("send to close collector")
)
