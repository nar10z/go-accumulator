/*
 * Copyright (c) 2023.
 *
 * License MIT (https://raw.githubusercontent.com/nar10z/go-accumulator/main/LICENSE)
 *
 * Developed thanks to Nikita Terentyev (nar10z). Use it for good, and let your code work without problems!
 */

package storage

import (
	"fmt"
	"strconv"
	"sync/atomic"
	"testing"
)

func BenchmarkStorage(b *testing.B) {
	const size = 10000

	type A struct {
		s string
		i int
	}

	b.ResetTimer()
	b.Run("#1. list", func(b *testing.B) {
		stor := NewStorageList[*A](size)
		sum := atomic.Int32{}

		for i := 0; i < b.N; i++ {
			s := strconv.FormatInt(int64(i), 10)

			if !stor.Put(&A{s: s, i: i}) {
				stor.Iterate(func(ee *A) {
					sum.Add(1)
				})
				stor.Clear()
			}
		}

		sum.Add(int32(stor.Len()))
		if sum.Load() != int32(b.N) {
			fmt.Printf("got=%d, want=%d\n", sum.Load(), b.N)
			b.Fail()
		}
	})
	b.Run("#2. gods/list", func(b *testing.B) {
		stor := NewStorageSinglyList[*A](size)
		sum := atomic.Int32{}

		for i := 0; i < b.N; i++ {
			s := strconv.FormatInt(int64(i), 10)

			if !stor.Put(&A{s: s, i: i}) {
				stor.Iterate(func(ee *A) {
					sum.Add(1)
				})
				stor.Clear()
			}
		}

		sum.Add(int32(stor.Len()))
		if sum.Load() != int32(b.N) {
			fmt.Printf("got=%d, want=%d\n", sum.Load(), b.N)
			b.Fail()
		}
	})
	b.Run("#3. slice", func(b *testing.B) {
		stor := NewStorageSlice[*A](size)
		sum := atomic.Int32{}

		for i := 0; i < b.N; i++ {
			s := strconv.FormatInt(int64(i), 10)

			if !stor.Put(&A{s: s, i: i}) {
				stor.Iterate(func(ee *A) {
					sum.Add(1)
				})
				stor.Clear()
			}
		}

		sum.Add(int32(stor.Len()))
		if sum.Load() != int32(b.N) {
			fmt.Printf("got=%d, want=%d\n", sum.Load(), b.N)
			b.Fail()
		}
	})
}
