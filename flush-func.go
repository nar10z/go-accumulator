/*
 * Copyright (c) 2023.
 *
 * License MIT (https://raw.githubusercontent.com/nar10z/go-accumulator/main/LICENSE)
 *
 * Developed thanks to Nikita Terentyev (nar10z). Use it for good, and let your code work without problems!
 */

package goaccum

import "context"

// FlushExec is called when the accumulated batch needs to be processed.
// The events slice is only valid for the duration of the function's execution.
// Do not retain a reference to the slice outside the function call.
type FlushExec[T any] func(ctx context.Context, events []T) error

func noop[T any](_ context.Context, _ []T) error {
	return nil
}
