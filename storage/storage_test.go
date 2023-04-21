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
	"testing"
)

func BenchmarkStorage(b *testing.B) {
	const size = 10000

	type A struct {
		s string
		i int
	}

	b.ResetTimer()
	b.Run("#1. Channel", func(b *testing.B) {
		stor := NewStorageChannel[*A](size)
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
	b.Run("#2. container/list", func(b *testing.B) {
		stor := NewStorageList[*A](size)
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
	b.Run("#3. gods/singlylinkedlist", func(b *testing.B) {
		stor := NewStorageSinglyList[*A](size)
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
	b.Run("#4. slice", func(b *testing.B) {
		stor := NewStorageSlice[*A](size)
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
