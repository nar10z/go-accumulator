/*
 * Copyright (c) 2023.
 *
 * License MIT (https://raw.githubusercontent.com/nar10z/go-accumulator/main/LICENSE)
 *
 * Developed thanks to Nikita Terentyev (nar10z). Use it for good, and let your code work without problems!
 */

package go_accumulator

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/sync/errgroup"
)

func Test_New(t *testing.T) {
	t.Parallel()

	t.Run("#1. Failed", func(t *testing.T) {
		coll, err := New[int](0, 0, nil)

		require.Error(t, err)
		assert.Nil(t, coll)
	})
	t.Run("#2. Success", func(t *testing.T) {
		coll, err := New[int](10, time.Millisecond*20, func(events []int) error { return nil })

		require.NoError(t, err)
		require.NotNil(t, coll)

		coll.Stop()
		col, ok := coll.(*accumulator[int])
		require.True(t, ok)
		require.NotNil(t, col)
		require.True(t, col.isClose.Load())
	})
}

func Test_accumulator(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	t.Run("#1.1. Only async", func(t *testing.T) {

		var (
			countWriters    = 2
			countAsyncEvent = 113
			summary         = 0
		)

		coll, err := New(100, time.Millisecond*50, func(events []int) error {
			summary += len(events)
			return nil
		})

		require.NoError(t, err)
		require.NotNil(t, coll)

		var wgEvents sync.WaitGroup

		for i := 0; i < countWriters; i++ {
			wgEvents.Add(1)
			go func() {
				defer wgEvents.Done()

				for i := 0; i < countAsyncEvent; i++ {
					require.NoError(t, coll.AddAsync(ctx, i))
				}
			}()
		}

		wgEvents.Wait()

		coll.Stop()

		require.Equal(t, countAsyncEvent*countWriters, summary)
	})
	t.Run("#1.2. Only sync", func(t *testing.T) {

		var (
			countSyncEvent = 3851
			summary        = 0
		)

		coll, err := New(100, time.Millisecond*100, func(events []int) error {
			summary += len(events)
			return nil
		})

		require.NoError(t, err)
		require.NotNil(t, coll)

		var errGr errgroup.Group
		errGr.SetLimit(5000)

		for i := 0; i < countSyncEvent; i++ {
			errGr.Go(func() error {
				return coll.AddSync(ctx, 1)
			})
		}
		require.NoError(t, errGr.Wait())

		coll.Stop()

		require.Equal(t, countSyncEvent, summary)
	})
	t.Run("#1.3. Async and sync", func(t *testing.T) {

		var (
			countSyncEvent  = 2454
			countAsyncEvent = 3913
			summary         = 0
		)

		coll, err := New(1000, time.Millisecond*100, func(events []int) error {
			summary += len(events)
			return nil
		})

		require.NoError(t, err)
		require.NotNil(t, coll)

		var wgEvents sync.WaitGroup

		wgEvents.Add(1)
		go func() {
			defer wgEvents.Done()

			for i := 0; i < countAsyncEvent; i++ {
				require.NoError(t, coll.AddAsync(ctx, i))
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
					return coll.AddSync(ctx, i)
				})
			}
			require.NoError(t, errGr.Wait())
		}()

		wgEvents.Wait()
		coll.Stop()

		require.Equal(t, countSyncEvent+countAsyncEvent, summary)
	})

	t.Run("#2.1. Long interval", func(t *testing.T) {

		var (
			countSyncEvent  = 1200
			countAsyncEvent = 6300
			summary         = 0
		)

		coll, err := New(1000, time.Minute*10, func(events []int) error {
			summary += len(events)
			return nil
		})

		require.NoError(t, err)
		require.NotNil(t, coll)

		var wgEvents sync.WaitGroup

		wgEvents.Add(1)
		go func() {
			defer wgEvents.Done()

			for i := 0; i < countAsyncEvent; i++ {
				require.NoError(t, coll.AddAsync(ctx, i))
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
					return coll.AddSync(ctx, i)
				})
			}
			require.NoError(t, errGr.Wait())
		}()

		wgEvents.Wait()
		coll.Stop()

		require.Equal(t, countSyncEvent+countAsyncEvent, summary)
	})
	t.Run("#2.2. Big size", func(t *testing.T) {

		var (
			countSyncEvent  = 1200
			countAsyncEvent = 6300
			summary         = 0
		)

		coll, err := New(1000000, time.Millisecond*50, func(events []int) error {
			summary += len(events)
			return nil
		})

		require.NoError(t, err)
		require.NotNil(t, coll)

		var wgEvents sync.WaitGroup

		wgEvents.Add(1)
		go func() {
			defer wgEvents.Done()

			for i := 0; i < countAsyncEvent; i++ {
				require.NoError(t, coll.AddAsync(ctx, i))
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
					return coll.AddSync(ctx, i)
				})
			}
			require.NoError(t, errGr.Wait())
		}()

		wgEvents.Wait()
		coll.Stop()

		require.Equal(t, countSyncEvent+countAsyncEvent, summary)
	})
	t.Run("#2.3. Context deadline", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(ctx, time.Microsecond)
		defer cancel()

		var (
			countAsyncEvent = 100
			summary         = 0
		)

		coll, err := New(1000, time.Millisecond*100, func(events []int) error {
			summary += len(events)
			return nil
		})

		require.NoError(t, err)
		require.NotNil(t, coll)

		time.Sleep(10 * time.Microsecond)

		for i := 0; i < countAsyncEvent; i++ {
			require.Error(t, coll.AddAsync(ctx, i))
		}

		coll.Stop()

		require.Equal(t, 0, summary)
	})
	t.Run("#2.4. Context deadline", func(t *testing.T) {
		ctxIn, cancelIn := context.WithTimeout(ctx, time.Nanosecond)
		defer cancelIn()

		var (
			countSyncEvent = 110
			summary        = 0
		)

		coll, err := New(100, time.Millisecond*100, func(events []int) error {
			summary += len(events)
			return nil
		})

		require.NoError(t, err)
		require.NotNil(t, coll)

		time.Sleep(10 * time.Microsecond)

		var errGr errgroup.Group
		errGr.SetLimit(50)
		for i := 0; i < countSyncEvent; i++ {
			i := i
			errGr.Go(func() error {
				return coll.AddSync(ctxIn, i)
			})
		}
		_ = errGr.Wait()

		coll.Stop()

		require.Equal(t, 0, summary)
	})
	t.Run("#2.5. Send to close buffer", func(t *testing.T) {
		var (
			countSyncEvent  = 30
			countAsyncEvent = 10
			summary         = 0
		)

		coll, err := New(1000, time.Millisecond*100, func(events []int) error {
			summary += len(events)
			return nil
		})
		coll.Stop()

		require.NoError(t, err)
		require.NotNil(t, coll)

		var wgEvents sync.WaitGroup

		wgEvents.Add(1)
		go func() {
			defer wgEvents.Done()

			for i := 0; i < countAsyncEvent; i++ {
				require.Error(t, coll.AddAsync(ctx, i))
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
					return coll.AddSync(ctx, i)
				})
			}
			require.Error(t, errGr.Wait())
		}()

		wgEvents.Wait()

		require.Equal(t, 0, summary)
	})
}
