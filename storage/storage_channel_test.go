/*
 * Copyright (c) 2023.
 *
 * Developed thanks to Nikita Terentyev (nar10z). Use it for good, and let your code work without problems!
 */

package storage

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_newEventStorage(t *testing.T) {
	t.Parallel()

	t.Run("#1. One", func(t *testing.T) {
		t.Parallel()

		stor := NewStorageList[int](10)
		allowed := stor.Put(1)

		require.True(t, allowed)

		data := stor.Get()
		require.NotEmpty(t, data)
		require.Len(t, data, 1)
	})
	t.Run("#2. More", func(t *testing.T) {
		t.Parallel()

		const (
			size = 1000
			n    = 1_000_000
		)
		stor := NewStorageList[int](size)
		sum := 0

		for i := 0; i < n; i++ {
			if !stor.Put(i) {
				sum += len(stor.Get())
			}
		}
		sum += len(stor.Get())

		require.Equal(t, sum, n)
	})
}
