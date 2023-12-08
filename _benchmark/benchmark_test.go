package accumulator_example

import (
	"context"
	"testing"
	"time"

	acc "github.com/lrweck/accumulator"
	goaccum "github.com/nar10z/go-accumulator"
	"golang.org/x/sync/errgroup"
)

const (
	flushSize     = 1000
	flushInterval = time.Second
)

type Data struct {
	i int
}

func Benchmark_accum(b *testing.B) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	b.Run("go-accumulator, async", func(b *testing.B) {
		summary := 0

		accumulator, _ := goaccum.New[Data](flushSize, flushInterval, func(events []Data) error {
			summary += len(events)
			return nil
		})

		for i := 0; i < b.N; i++ {
			_ = accumulator.AddAsync(ctx, Data{i: i})
		}

		accumulator.Stop()

		if summary != b.N {
			b.Fail()
		}
	})
	b.Run("lrweck/accumulator", func(b *testing.B) {
		summary := 0

		inputChan := make(chan Data, flushSize)
		batch := acc.New(inputChan, flushSize, flushInterval)

		go func() {
			for i := 0; i < b.N; i++ {
				inputChan <- Data{i: i}
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

	b.Run("go-accumulator, sync", func(b *testing.B) {
		summary := 0

		accumulator, _ := goaccum.New[Data](flushSize, flushInterval, func(events []Data) error {
			summary += len(events)
			return nil
		})

		var errGr errgroup.Group
		errGr.SetLimit(flushSize)
		for i := 0; i < b.N; i++ {
			errGr.Go(func() error {
				return accumulator.AddSync(ctx, Data{i: i})
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
		n1 := b.N / 2
		n2 := b.N - n1

		accumulator, _ := goaccum.New[Data](flushSize, flushInterval, func(events []Data) error {
			summary += len(events)
			return nil
		})

		for i := 0; i < n1; i++ {
			_ = accumulator.AddAsync(ctx, Data{i: i})
		}

		var errGr errgroup.Group
		errGr.SetLimit(flushSize)
		for i := 0; i < n2; i++ {
			errGr.Go(func() error {
				return accumulator.AddSync(ctx, Data{i: i})
			})
		}

		_ = errGr.Wait()
		accumulator.Stop()

		if summary != b.N {
			b.Fail()
		}
	})
}
