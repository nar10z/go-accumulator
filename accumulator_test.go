/*
 * Copyright (c) 2023.
 *
 * License MIT (https://raw.githubusercontent.com/nar10z/go-accumulator/main/LICENSE)
 *
 * Developed thanks to Nikita Terentyev (nar10z). Use it for good, and let your code work without problems!
 */

package goaccum

import (
	"context"
	"errors"
	"runtime"
	"slices"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/sync/errgroup"
)

func Test_New_Success(t *testing.T) {
	t.Parallel()

	coll := New[int](10, time.Millisecond, time.Second, func(_ context.Context, events []int) error { return nil })
	require.NotNil(t, coll)

	coll.Stop()
	assert.True(t, coll.IsClosed())
}

func Test_New_EmptyParams(t *testing.T) {
	t.Parallel()

	coll := New[int](0, 0, 0, nil)
	require.NotNil(t, coll)

	assert.NotEmpty(t, coll.flushFunc)

	var (
		wg         sync.WaitGroup
		countStops = 100
	)

	wg.Add(countStops)
	for range countStops {
		go func() {
			coll.Stop()
			assert.True(t, coll.IsClosed())
			wg.Done()
		}()
	}

	wg.Wait()
}

func Test_New_NegativeIntervalUsesDefault(t *testing.T) {
	t.Parallel()

	coll := New[int](10, -time.Second, time.Second, nil)
	require.NotNil(t, coll)
	coll.Stop()
}

func Test_Accumulator_OnlyAsync(t *testing.T) {
	var (
		countWriters    = 2
		countAsyncEvent = 113
		summary         = 0
	)

	coll := New(100, time.Millisecond*50, time.Millisecond*10, func(_ context.Context, events []int) error {
		summary += len(events)
		return nil
	})
	require.NotNil(t, coll)

	var wgEvents sync.WaitGroup
	for range countWriters {
		wgEvents.Go(func() {
			for i := range countAsyncEvent {
				require.NoError(t, coll.AddAsync(t.Context(), i))
			}
		})
	}
	wgEvents.Wait()
	require.False(t, coll.IsClosed())
	coll.Stop()
	require.Equal(t, countAsyncEvent*countWriters, summary)
	require.True(t, coll.IsClosed())
}

func Test_Accumulator_OnlySync(t *testing.T) {
	var (
		countSyncEvent = 3851
		summary        = 0
	)
	coll := New(100, time.Millisecond*100, time.Millisecond*10, func(_ context.Context, events []int) error {
		summary += len(events)
		return nil
	})
	require.NotNil(t, coll)
	var errGr errgroup.Group
	errGr.SetLimit(5000)
	for range countSyncEvent {
		errGr.Go(func() error { return coll.AddSync(t.Context(), 1) })
	}
	require.NoError(t, errGr.Wait())
	coll.Stop()
	require.Equal(t, countSyncEvent, summary)
}

func Test_Accumulator_AsyncAndSync(t *testing.T) {
	var (
		countSyncEvent  = 2454
		countAsyncEvent = 3913
		summary         = 0
	)

	coll := New(1000, time.Millisecond*100, time.Millisecond*10, func(_ context.Context, events []int) error {
		summary += len(events)
		return nil
	})
	require.NotNil(t, coll)

	var wgEvents sync.WaitGroup
	wgEvents.Go(func() {
		for i := range countAsyncEvent {
			require.NoError(t, coll.AddAsync(t.Context(), i))
		}
	})
	wgEvents.Go(func() {
		var errGr errgroup.Group
		errGr.SetLimit(5000)
		for i := range countSyncEvent {
			errGr.Go(func() error { return coll.AddSync(t.Context(), i) })
		}
		require.NoError(t, errGr.Wait())
	})

	wgEvents.Wait()
	coll.Stop()
	require.Equal(t, countSyncEvent+countAsyncEvent, summary)
}

func Test_Accumulator_LongInterval(t *testing.T) {
	var (
		countSyncEvent  = 1200
		countAsyncEvent = 6300
		summary         = 0
	)

	coll := New(1000, time.Minute*10, time.Millisecond*10, func(_ context.Context, events []int) error {
		summary += len(events)
		return nil
	})
	require.NotNil(t, coll)

	var wgEvents sync.WaitGroup
	wgEvents.Go(func() {
		for i := range countAsyncEvent {
			require.NoError(t, coll.AddAsync(t.Context(), i))
		}
	})
	wgEvents.Go(func() {
		var errGr errgroup.Group
		errGr.SetLimit(5000)
		for i := range countSyncEvent {
			errGr.Go(func() error { return coll.AddSync(t.Context(), i) })
		}
		require.NoError(t, errGr.Wait())
	})
	wgEvents.Wait()
	coll.Stop()
	require.Equal(t, countSyncEvent+countAsyncEvent, summary)
}

func Test_Accumulator_BigSize(t *testing.T) {
	var (
		countSyncEvent  = 1200
		countAsyncEvent = 6300
		summary         = 0
	)

	coll := New(1000000, time.Millisecond*50, time.Millisecond*10, func(_ context.Context, events []int) error {
		summary += len(events)
		return nil
	})
	require.NotNil(t, coll)

	var wgEvents sync.WaitGroup
	wgEvents.Go(func() {
		for i := range countAsyncEvent {
			require.NoError(t, coll.AddAsync(t.Context(), i))
		}
	})
	wgEvents.Go(func() {
		var errGr errgroup.Group
		errGr.SetLimit(5000)
		for i := range countSyncEvent {
			errGr.Go(func() error { return coll.AddSync(t.Context(), i) })
		}
		require.NoError(t, errGr.Wait())
	})
	wgEvents.Wait()

	coll.Stop()
	require.Equal(t, countSyncEvent+countAsyncEvent, summary)
}

func Test_Accumulator_ContextDeadlineAsync(t *testing.T) {
	ctxIn, cancelIn := context.WithTimeout(t.Context(), time.Nanosecond)
	defer cancelIn()

	var (
		countAsyncEvent = 100
		summary         = 0
	)

	coll := New(1000, time.Millisecond*100, time.Millisecond*10, func(_ context.Context, events []int) error {
		summary += len(events)
		return nil
	})
	require.NotNil(t, coll)

	time.Sleep(time.Second)

	for i := range countAsyncEvent {
		require.Error(t, coll.AddAsync(ctxIn, i))
	}

	coll.Stop()
	require.Equal(t, 0, summary)
}

func Test_Accumulator_ContextDeadlineSync(t *testing.T) {
	ctxIn, cancelIn := context.WithTimeout(t.Context(), time.Nanosecond)
	defer cancelIn()

	var (
		countSyncEvent = 110
		summary        = 0
	)

	coll := New(100, time.Millisecond*100, time.Millisecond*10, func(_ context.Context, events []int) error {
		summary += len(events)
		return nil
	})
	require.NotNil(t, coll)

	time.Sleep(time.Second)

	var errGr errgroup.Group
	errGr.SetLimit(50)
	for i := range countSyncEvent {
		errGr.Go(func() error { return coll.AddSync(ctxIn, i) })
	}
	_ = errGr.Wait()

	coll.Stop()

	require.Equal(t, 0, summary)
}

func Test_Accumulator_SendToCloseBuffer(t *testing.T) {
	var (
		countSyncEvent  = 30
		countAsyncEvent = 10
		summary         = 0
	)

	coll := New(1000, time.Millisecond*100, time.Millisecond*10, func(_ context.Context, events []int) error {
		summary += len(events)
		return nil
	})
	coll.Stop()
	require.NotNil(t, coll)

	var wgEvents sync.WaitGroup
	wgEvents.Go(func() {
		for i := range countAsyncEvent {
			require.Error(t, coll.AddAsync(t.Context(), i))
		}
	})
	wgEvents.Go(func() {
		var errGr errgroup.Group
		errGr.SetLimit(5000)
		for i := range countSyncEvent {
			errGr.Go(func() error { return coll.AddSync(t.Context(), i) })
		}
		require.Error(t, errGr.Wait())
	})

	wgEvents.Wait()
	require.Equal(t, 0, summary)
}

func Test_Accumulator_ReturnsFlushErrorInAddSync(t *testing.T) {
	wantErr := errors.New("some")

	coll := New(2, time.Millisecond, time.Millisecond, func(_ context.Context, events []int) error { return wantErr })
	require.NotNil(t, coll)

	errAdd := coll.AddSync(t.Context(), 1)
	require.ErrorIs(t, errAdd, wantErr)

	coll.Stop()
}

func Test_Accumulator_EqualResult(t *testing.T) {
	var (
		result []int
		want   = []int{0, 1, 2, 3, 4}
	)

	coll := New(2, time.Millisecond*10, time.Millisecond*10, func(_ context.Context, events []int) error {
		result = append(result, events...)
		return nil
	})
	require.NotNil(t, coll)

	var errGr errgroup.Group
	errGr.SetLimit(5)
	for i := range 5 {
		errGr.Go(func() error { return coll.AddSync(t.Context(), i) })
	}

	require.NoError(t, errGr.Wait())

	coll.Stop()

	slices.Sort(result)
	require.Equal(t, result, want)
}

// Test_ConcurrentStop_Idempotent verifies that multiple calls to Stop() are safe.
func Test_ConcurrentStop_Idempotent(t *testing.T) {
	t.Parallel()

	coll := New[int](
		10,
		time.Millisecond*10,
		time.Second,
		func(_ context.Context, events []int) error { return nil },
	)

	var wg sync.WaitGroup
	for range 20 {
		wg.Go(func() {
			coll.Stop()
		})
	}

	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		// OK
	case <-time.After(time.Second * 2):
		t.Fatal("concurrent Stop() calls deadlocked")
	}

	require.True(t, coll.IsClosed())
}

// Test_AddAsync_RespectsContext checks that AddAsync is actually canceled in the context
// and does not block indefinitely on a full channel.
func Test_AddAsync_RespectsContext(t *testing.T) {
	t.Parallel()

	const flushSize = 3

	// Создаём аккумулятор с flushFunc, которая НИЧЕГО не делает (никогда не flush'ит)
	// и маленьким буфером канала, чтобы быстро забить его.
	coll := New[int](
		flushSize, // flushSize
		time.Hour, // flushInterval — очень долгий
		time.Second,
		nil,
	)

	// Забиваем канал
	for i := range flushSize {
		require.NoError(t, coll.AddAsync(t.Context(), i))
	}

	// Теперь канал забит. Пробуем добавить с отменённым контекстом.
	ctx, cancel := context.WithCancel(t.Context())
	cancel()

	err := coll.AddAsync(ctx, 999)
	require.Error(t, err)
	require.ErrorIs(t, err, context.Canceled)

	coll.Stop()
}

// Test_NoGoroutineLeak checks that no accumulator routines remain after Stop().
func Test_NoGoroutineLeak(t *testing.T) {
	// Не t.Parallel() — используем runtime.ReadGoroutineStats

	before := runtime.NumGoroutine()

	for range 50 {
		coll := New[int](
			10,
			time.Millisecond*10,
			time.Second,
			nil,
		)

		for j := range 5 {
			_ = coll.AddAsync(t.Context(), j)
		}

		coll.Stop()
	}

	// Даём время на завершение горутин
	time.Sleep(time.Millisecond * 200)
	runtime.GC()
	time.Sleep(time.Millisecond * 100)

	after := runtime.NumGoroutine()
	// Допускаем небольшую погрешность
	require.LessOrEqual(t, after, before+2, "goroutine leak detected: before=%d, after=%d", before, after)
}

// Test_Stop_DoesNotPanicOnConcurrentAdd checks that Stop() is safe when AddAsync and AddSync are called concurrently.
func Test_Stop_DoesNotPanicOnConcurrentAdd(t *testing.T) {
	t.Parallel()

	const flushSize = 12

	coll := New[int](
		flushSize,
		time.Millisecond*50,
		time.Second,
		nil,
	)

	var wg sync.WaitGroup
	for range flushSize {
		wg.Go(func() {
			for j := range 50 {
				_ = coll.AddAsync(t.Context(), j)
			}
		})
	}

	// Конкурентно вызываем Stop
	go coll.Stop()

	wg.Wait()
	// Если дошли сюда без паники — тест пройден
}

// ---------------------------------------------------------------------------
// Additional tests for uncovered error/corner branches
// ---------------------------------------------------------------------------

func Test_AddAsync_RecoverFromSendToClosedChannel(t *testing.T) {
	t.Parallel()

	coll := &Accumulator[int]{
		chEvents: make(chan eventExtended[int]),
		chDone:   make(chan struct{}),
	}
	close(coll.chEvents)

	err := coll.AddAsync(t.Context(), 1)
	require.Error(t, err)
	require.Contains(t, err.Error(), "recover")
}

func Test_AddSync_RecoverFromSendToClosedChannel(t *testing.T) {
	t.Parallel()

	coll := &Accumulator[int]{
		chEvents: make(chan eventExtended[int]),
		chDone:   make(chan struct{}),
	}
	close(coll.chEvents)

	err := coll.AddSync(t.Context(), 1)
	require.Error(t, err)
	require.Contains(t, err.Error(), "recover")
}

func Test_AddSync_ContextCancelledWhileWaitingFallback(t *testing.T) {
	t.Parallel()

	coll := &Accumulator[int]{
		chEvents: make(chan eventExtended[int], 1),
		chDone:   make(chan struct{}),
	}

	ctx, cancel := context.WithCancel(t.Context())
	defer cancel()

	done := make(chan error, 1)
	go func() {
		done <- coll.AddSync(ctx, 42)
	}()

	evt := <-coll.chEvents
	require.NotNil(t, evt.fallback)

	cancel()

	select {
	case err := <-done:
		require.Error(t, err)
		require.ErrorIs(t, err, context.Canceled)
	case <-time.After(time.Second):
		t.Fatal("AddSync did not return after context cancellation")
	}
}

func Test_AddSync_PropagatesFlushErrorToAllWaiters(t *testing.T) {
	t.Parallel()

	wantErr := errors.New("flush failed")
	const count = 20

	coll := New[int](5, time.Hour, time.Second, func(_ context.Context, _ []int) error {
		return wantErr
	})
	defer coll.Stop()

	var (
		errCount atomic.Int64
		errGr    errgroup.Group
	)
	errGr.SetLimit(10)

	for i := range count {
		errGr.Go(func() error {
			err := coll.AddSync(t.Context(), i)
			if errors.Is(err, wantErr) {
				errCount.Add(1)
			}
			return nil
		})
	}
	require.NoError(t, errGr.Wait())
	require.Equal(t, int64(count), errCount.Load())
}

func Test_AddAsync_BlocksUntilSpaceAvailable(t *testing.T) {
	t.Parallel()

	coll := &Accumulator[int]{
		chEvents: make(chan eventExtended[int], 1),
		chDone:   make(chan struct{}),
	}
	coll.chEvents <- eventExtended[int]{e: 1} // fill buffer

	ctx, cancel := context.WithTimeout(t.Context(), 500*time.Millisecond)
	defer cancel()

	done := make(chan error, 1)
	start := time.Now()
	go func() {
		done <- coll.AddAsync(ctx, 2)
	}()

	time.Sleep(80 * time.Millisecond)
	select {
	case err := <-done:
		t.Fatalf("AddAsync returned too early: %v", err)
	default:
	}

	<-coll.chEvents // free one slot

	select {
	case err := <-done:
		require.NoError(t, err)
	case <-time.After(time.Second):
		t.Fatal("AddAsync did not complete after freeing channel space")
	}

	require.GreaterOrEqual(t, time.Since(start), 80*time.Millisecond)
}

func Test_New_DefaultsFlushTimeoutWhenNonPositive(t *testing.T) {
	t.Parallel()

	interval := 25 * time.Millisecond
	coll := New[int](10, interval, -time.Second, nil)
	defer coll.Stop()

	require.Equal(t, interval, coll.flushTimeout)
}
