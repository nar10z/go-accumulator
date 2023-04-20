/*
 * Copyright (c) 2023.
 *
 * Developed thanks to Nikita Terentyev (nar10z). Use it for good, and let your code work without problems!
 */

package go_events_accumulator

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/sync/errgroup"
)

func Test_NewAccumulator(t *testing.T) {
	t.Parallel()

	t.Run("#1. Failed", func(t *testing.T) {
		acc, err := NewAccumulator[int](0, 0, nil)

		require.Error(t, err)
		assert.Nil(t, acc)
	})
	t.Run("#2. Success", func(t *testing.T) {
		acc, err := NewAccumulator[int](10, time.Millisecond*20, func(events []int) error { return nil })

		require.NoError(t, err)
		require.NotNil(t, acc)

		acc.Stop()
		require.True(t, acc.isClose.Load())
	})
}

func Test_Accumulator(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	t.Run("#1.1. Only async", func(t *testing.T) {

		var (
			countWriters    = 2
			countAsyncEvent = 113
			summary         = 0
		)

		acc, err := NewAccumulator(100, time.Millisecond*50, func(events []int) error {
			time.Sleep(time.Millisecond)
			summary += len(events)
			return nil
		})

		require.NoError(t, err)
		require.NotNil(t, acc)

		var wgEvents sync.WaitGroup

		for i := 0; i < countWriters; i++ {
			wgEvents.Add(1)
			go func() {
				defer wgEvents.Done()

				for i := 0; i < countAsyncEvent; i++ {
					require.NoError(t, acc.AddAsync(ctx, i))
				}
			}()
		}

		wgEvents.Wait()

		acc.Stop()

		require.Equal(t, countAsyncEvent*countWriters, summary)
	})
	t.Run("#1.2. Only sync", func(t *testing.T) {

		var (
			countSyncEvent = 3851
			summary        = 0
		)

		acc, err := NewAccumulator(100, time.Millisecond*100, func(events []int) error {
			time.Sleep(time.Millisecond)
			summary += len(events)
			return nil
		})

		require.NoError(t, err)
		require.NotNil(t, acc)

		var errGr errgroup.Group
		errGr.SetLimit(5000)

		for i := 0; i < countSyncEvent; i++ {
			i := i
			errGr.Go(func() error {
				return acc.AddSync(ctx, i)
			})
		}
		require.NoError(t, errGr.Wait())

		acc.Stop()

		require.Equal(t, countSyncEvent, summary)
	})
	t.Run("#1.3. Async and sync", func(t *testing.T) {

		var (
			countSyncEvent  = 2454
			countAsyncEvent = 3913
			summary         = 0
		)

		acc, err := NewAccumulator(1000, time.Millisecond*100, func(events []int) error {
			time.Sleep(time.Millisecond)
			summary += len(events)
			return nil
		})

		require.NoError(t, err)
		require.NotNil(t, acc)

		var wgEvents sync.WaitGroup

		wgEvents.Add(1)
		go func() {
			defer wgEvents.Done()

			for i := 0; i < countAsyncEvent; i++ {
				require.NoError(t, acc.AddAsync(ctx, i))
			}
		}()

		wgEvents.Add(1)
		go func() {
			defer wgEvents.Done()

			var errGr errgroup.Group
			errGr.SetLimit(5000)

			for i := 0; i < countSyncEvent; i++ {
				i := i
				errGr.Go(func() error {
					return acc.AddSync(ctx, i)
				})
			}
			require.NoError(t, errGr.Wait())
		}()

		wgEvents.Wait()
		acc.Stop()

		require.Equal(t, countSyncEvent+countAsyncEvent, summary)
	})

	t.Run("#2.1. Long interval", func(t *testing.T) {

		var (
			countSyncEvent  = 1200
			countAsyncEvent = 6300
			summary         = 0
		)

		acc, err := NewAccumulator(1000, time.Minute*10, func(events []int) error {
			time.Sleep(time.Millisecond)
			summary += len(events)
			return nil
		})

		require.NoError(t, err)
		require.NotNil(t, acc)

		var wgEvents sync.WaitGroup

		wgEvents.Add(1)
		go func() {
			defer wgEvents.Done()

			for i := 0; i < countAsyncEvent; i++ {
				require.NoError(t, acc.AddAsync(ctx, i))
			}
		}()

		wgEvents.Add(1)
		go func() {
			defer wgEvents.Done()

			var errGr errgroup.Group
			errGr.SetLimit(5000)

			for i := 0; i < countSyncEvent; i++ {
				i := i
				errGr.Go(func() error {
					return acc.AddSync(ctx, i)
				})
			}
			require.NoError(t, errGr.Wait())
		}()

		wgEvents.Wait()
		acc.Stop()

		require.Equal(t, countSyncEvent+countAsyncEvent, summary)
	})
	t.Run("#2.2. Big size", func(t *testing.T) {

		var (
			countSyncEvent  = 1200
			countAsyncEvent = 6300
			summary         = 0
		)

		acc, err := NewAccumulator(1000000, time.Millisecond*50, func(events []int) error {
			time.Sleep(time.Millisecond)
			summary += len(events)
			return nil
		})

		require.NoError(t, err)
		require.NotNil(t, acc)

		var wgEvents sync.WaitGroup

		wgEvents.Add(1)
		go func() {
			defer wgEvents.Done()

			for i := 0; i < countAsyncEvent; i++ {
				require.NoError(t, acc.AddAsync(ctx, i))
			}
		}()

		wgEvents.Add(1)
		go func() {
			defer wgEvents.Done()

			var errGr errgroup.Group
			errGr.SetLimit(5000)

			for i := 0; i < countSyncEvent; i++ {
				i := i
				errGr.Go(func() error {
					return acc.AddSync(ctx, i)
				})
			}
			require.NoError(t, errGr.Wait())
		}()

		wgEvents.Wait()
		acc.Stop()

		require.Equal(t, countSyncEvent+countAsyncEvent, summary)
	})
	t.Run("#2.3. Context deadline", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(ctx, time.Microsecond)
		defer cancel()

		var (
			countSyncEvent  = 300
			countAsyncEvent = 100
			summary         = 0
		)

		acc, err := NewAccumulator(1000, time.Millisecond*100, func(events []int) error {
			summary += len(events)
			return nil
		})

		require.NoError(t, err)
		require.NotNil(t, acc)

		time.Sleep(time.Second)

		var wgEvents sync.WaitGroup

		wgEvents.Add(1)
		go func() {
			defer wgEvents.Done()

			for i := 0; i < countAsyncEvent; i++ {
				require.Error(t, acc.AddAsync(ctx, i))
			}
		}()

		wgEvents.Add(1)
		go func() {
			defer wgEvents.Done()

			var errGr errgroup.Group
			errGr.SetLimit(5000)

			for i := 0; i < countSyncEvent; i++ {
				i := i
				errGr.Go(func() error {
					return acc.AddSync(ctx, i)
				})
			}
			require.Error(t, errGr.Wait())
		}()

		wgEvents.Wait()
		acc.Stop()

		require.Equal(t, 0, summary)
	})
	t.Run("#2.4. Send to close buffer", func(t *testing.T) {
		var (
			countSyncEvent  = 30
			countAsyncEvent = 10
			summary         = 0
		)

		acc, err := NewAccumulator(1000, time.Millisecond*100, func(events []int) error {
			summary += len(events)
			return nil
		})
		acc.Stop()

		require.NoError(t, err)
		require.NotNil(t, acc)

		var wgEvents sync.WaitGroup

		wgEvents.Add(1)
		go func() {
			defer wgEvents.Done()

			for i := 0; i < countAsyncEvent; i++ {
				require.Error(t, acc.AddAsync(ctx, i))
			}
		}()

		wgEvents.Add(1)
		go func() {
			defer wgEvents.Done()

			var errGr errgroup.Group
			errGr.SetLimit(5000)

			for i := 0; i < countSyncEvent; i++ {
				i := i
				errGr.Go(func() error {
					return acc.AddSync(ctx, i)
				})
			}
			require.Error(t, errGr.Wait())
		}()

		wgEvents.Wait()

		require.Equal(t, 0, summary)
	})
}
