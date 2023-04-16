package go_events_accumulator

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestNewAccumulator(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	var (
		countSyncEvents = 20
		countAsyncEvent = 10
		summary         = 0
	)

	t.Run("#1", func(t *testing.T) {
		acc, err := NewAccumulator(Opts[int]{
			FlushSize:     10,
			FlushInterval: time.Second,
			FlushFunc: func(events []int) error {
				summary += len(events)
				return nil
			},
		})

		require.NoError(t, err)
		require.NotNil(t, acc)

		var wgEvents sync.WaitGroup

		wgEvents.Add(1)
		go func() {
			for i := 0; i < countAsyncEvent; i++ {
				err := acc.AddAsync(ctx, i)
				require.NoError(t, err)
			}
			wgEvents.Done()
		}()

		go func() {
			for i := 0; i < countSyncEvents; i++ {
				wgEvents.Add(1)

				i := i
				go func() {
					err := acc.AddSync(ctx, i)
					require.NoError(t, err)
					wgEvents.Done()
				}()
			}
		}()

		wgEvents.Wait()

		fmt.Printf("stop accumulator \n")
		acc.Stop()

		require.Equal(t, countSyncEvents+countAsyncEvent, summary)
	})
}
