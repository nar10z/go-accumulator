package accumulator_example

import (
	"context"
	"math/rand"
	"sync"
	"testing"
	"time"

	acc "github.com/lrweck/accumulator"
	goaccum "github.com/nar10z/go-accumulator"
	"golang.org/x/sync/errgroup"
)

const (
	flushSize     = 5000
	flushInterval = time.Millisecond * 50
	flushTimout   = time.Millisecond * 40
)

type Data struct {
	b bool
	i int
	f float32
}

func newData(i int) Data {
	return Data{
		b: rand.Intn(1) == 0,
		i: i,
		f: float32(i) * 1.07,
	}
}

func Benchmark_accum(b *testing.B) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	b.Run("go-accumulator async", func(b *testing.B) {
		summary := 0

		b.ResetTimer()

		accumulator := goaccum.New[Data](flushSize, flushInterval, flushTimout, func(_ context.Context, events []Data) error {
			summary += len(events)
			return nil
		})

		for i := 0; i < b.N; i++ {
			_ = accumulator.AddAsync(ctx, newData(i))
		}

		accumulator.Stop()

		if summary != b.N {
			b.Fail()
		}
	})

	b.Run("lrweck/accumulator", func(b *testing.B) {
		summary := 0
		inputChan := make(chan Data, flushSize)

		b.ResetTimer()

		batch := acc.New(inputChan, flushSize, flushInterval)

		go func() {
			for i := 0; i < b.N; i++ {
				inputChan <- newData(i)
			}
			close(inputChan)
		}()

		_ = batch.Accumulate(ctx, func(o acc.CallOrigin, items []Data) {
			summary += len(items)
		})

		if summary != b.N {
			b.Fail()
		}
	})

	b.Run("go-accumulator sync", func(b *testing.B) {
		summary := 0
		errGr := errgroup.Group{}

		errGr.SetLimit(flushSize * 1.5)

		b.ResetTimer()

		accumulator := goaccum.New[Data](flushSize, flushInterval, flushTimout, func(_ context.Context, events []Data) error {
			summary += len(events)
			return nil
		})

		for i := 0; i < b.N; i++ {
			errGr.Go(func() error {
				return accumulator.AddSync(ctx, newData(i))
			})
		}

		_ = errGr.Wait()
		accumulator.Stop()

		if summary != b.N {
			b.Fail()
		}
	})

	b.Run("go-accumulator", func(b *testing.B) {
		summary := 0
		n1 := rand.Intn(b.N)
		n2 := b.N - n1
		errGr := errgroup.Group{}
		wg := sync.WaitGroup{}

		errGr.SetLimit(flushSize * 1.5)

		b.ResetTimer()

		accumulator := goaccum.New[Data](flushSize, flushInterval, flushTimout, func(_ context.Context, events []Data) error {
			summary += len(events)
			return nil
		})

		wg.Add(1)
		go func() {
			for i := 0; i < n1; i++ {
				_ = accumulator.AddAsync(ctx, newData(i))
			}

			wg.Done()
		}()

		wg.Add(1)
		go func() {
			for i := 0; i < n2; i++ {
				errGr.Go(func() error {
					return accumulator.AddSync(ctx, newData(i))
				})
			}

			wg.Done()
		}()

		wg.Wait()

		_ = errGr.Wait()
		accumulator.Stop()

		if summary != b.N {
			b.Fail()
		}
	})
}
