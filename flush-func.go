/*
 * Copyright (c) 2023.
 *
 * License MIT (https://raw.githubusercontent.com/nar10z/go-accumulator/main/LICENSE)
 *
 * Developed thanks to Nikita Terentyev (nar10z). Use it for good, and let your code work without problems!
 */

package goaccum

import "context"

// FlushExec a function to call when an action needs to be performed
type FlushExec[T any] func(ctx context.Context, events []T) error

func noop[T any](_ context.Context, _ []T) error {
	return nil
}
