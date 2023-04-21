/*
 * Copyright (c) 2023.
 *
 * License MIT (https://raw.githubusercontent.com/nar10z/go-accumulator/main/LICENSE)
 *
 * Developed thanks to Nikita Terentyev (nar10z). Use it for good, and let your code work without problems!
 */

package main

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	gocoll "github.com/nar10z/go-accumulator"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	const (
		countSync  = 4
		countAsync = 3
	)

	accumulator, err := gocoll.New[string](3, time.Second, func(events []string) error {
		fmt.Printf("Start flush %d events:\n", len(events))
		for i, e := range events {
			fmt.Printf(" - %d) %s\n", i+1, e)
		}
		fmt.Printf("Finish\n")
		fmt.Printf(strings.Repeat("-", 100))
		fmt.Printf("\n")
		return nil
	})
	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup
	wg.Add(countSync + countAsync)

	go func() {
		for i := 0; i < countAsync; i++ {
			errE := accumulator.AddAsync(ctx, fmt.Sprintf("async №%d", i))
			if errE != nil {
				fmt.Printf("failed add event: %v\n", errE)
			}
			wg.Done()
		}
	}()

	go func() {
		for i := 0; i < countSync; i++ {
			i := i
			go func() {
				defer wg.Done()

				errE := accumulator.AddSync(ctx, fmt.Sprintf("sync №%d", i))
				if errE != nil {
					fmt.Printf("failed add event: %v\n", errE)
				}
			}()
		}
	}()

	wg.Wait()

	accumulator.Stop()
}
