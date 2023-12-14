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

	goaccum "github.com/nar10z/go-accumulator"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	const (
		countSync  = 4
		countAsync = 3
	)

	accumulator := goaccum.New[string](3, time.Second, func(events []string) error {
		fmt.Printf("Start flush %d events:\n", len(events))
		for _, e := range events {
			fmt.Printf(" - %s\n", e)
		}
		fmt.Printf("Finish\n%s\n", strings.Repeat("-", 20))
		return nil
	})

	var wg sync.WaitGroup
	wg.Add(countSync + countAsync)

	go func() {
		for i := 0; i < countAsync; i++ {
			err := accumulator.AddAsync(ctx, fmt.Sprintf("async #%d", i))
			if err != nil {
				fmt.Printf("failed add event: %v\n", err)
			}
			wg.Done()
		}
	}()

	go func() {
		for i := 0; i < countSync; i++ {
			i := i
			go func() {
				err := accumulator.AddSync(ctx, fmt.Sprintf("sync #%d", i))
				if err != nil {
					fmt.Printf("failed add event: %v\n", err)
				}
				wg.Done()
			}()
		}
	}()

	wg.Wait()

	accumulator.Stop()
}
