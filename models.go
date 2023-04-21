/*
 * Copyright (c) 2023.
 *
 * License MIT (https://raw.githubusercontent.com/nar10z/go-collector/main/LICENSE)
 *
 * Developed thanks to Nikita Terentyev (nar10z). Use it for good, and let your code work without problems!
 */

package go_collector

import "context"

// FlushExec a function to call when an action needs to be performed
type FlushExec[T comparable] func(events []T) error

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
