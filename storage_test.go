package go_events_accumulator

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_newEventStorage(t *testing.T) {
	t.Parallel()

	t.Run("#1. One", func(t *testing.T) {
		t.Parallel()

		stor := newEventStorage[int](10)
		l := stor.put(&eventExtend[int]{
			e: 1,
		})

		require.Equal(t, 1, l)

		data := stor.get()
		require.NotEmpty(t, data)
		require.Len(t, data, 1)
	})
	t.Run("#1. More", func(t *testing.T) {
		t.Parallel()

		const (
			size = 1000
			n    = 1_000_000
		)
		stor := newEventStorage[int](size)
		sum := 0

		for i := 0; i < n; i++ {
			if stor.put(&eventExtend[int]{e: i}) == size {
				sum += len(stor.get())
			}
		}
		sum += len(stor.get())

		require.Equal(t, sum, n)
	})
}

func BenchmarkStorage(b *testing.B) {

	const size = 100

	b.ResetTimer()

	b.Run("#1.", func(b *testing.B) {
		stor := newEventStorage[int](size)
		sum := 0

		for i := 0; i < b.N; i++ {
			if stor.put(&eventExtend[int]{e: i}) == size {
				sum += len(stor.get())
			}
		}

		sum += len(stor.get())
		if sum != b.N {
			b.Fail()
		}
	})
	b.Run("#2.", func(b *testing.B) {
		stor := newEventStorage[int](size)
		sum := 0

		for i := 0; i < b.N; i++ {
			l := stor.put(&eventExtend[int]{e: i})
			if l == size {
				sum += len(stor.get())
			}
		}

		sum += len(stor.get())
		if sum != b.N {
			b.Fail()
		}
	})
}
