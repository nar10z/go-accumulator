/*
 * Copyright (c) 2023.
 *
 * License MIT (https://raw.githubusercontent.com/nar10z/go-collector/main/LICENSE)
 *
 * Developed thanks to Nikita Terentyev (nar10z). Use it for good, and let your code work without problems!
 */

package go_collector

import "sync/atomic"

type eventExtended[T comparable] struct {
	// return error of flush operation
	fallback chan<- error
	// original data
	e T

	done atomic.Bool
}
