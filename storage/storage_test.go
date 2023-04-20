package storage

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_newEventStorage(t *testing.T) {
	t.Parallel()

	t.Run("#1. One", func(t *testing.T) {
		t.Parallel()

		stor := NewEventStorage[int](10)
		allowed := stor.Put(1)

		require.True(t, allowed)

		data := stor.Get()
		require.NotEmpty(t, data)
		require.Len(t, data, 1)
	})
	t.Run("#1. More", func(t *testing.T) {
		t.Parallel()

		const (
			size = 1000
			n    = 1_000_000
		)
		stor := NewEventStorage[int](size)
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

func BenchmarkStorage(b *testing.B) {
	const size = 100

	type A struct {
		s string
		i int
	}

	b.ResetTimer()
	b.Run("#1.", func(b *testing.B) {
		stor := NewEventStorage[*A](size)
		sum := 0

		for i := 0; i < b.N; i++ {
			s := strconv.FormatInt(int64(i), 10)

			if !stor.Put(&A{s: s, i: i}) {
				sum += len(stor.Get())
			}
		}

		sum += len(stor.Get())
		if sum != b.N {
			fmt.Printf("got=%d, want=%d\n", sum, b.N)
			b.Fail()
		}
	})
}
