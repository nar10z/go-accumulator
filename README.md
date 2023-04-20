# go-collector

[![Go Reference](https://pkg.go.dev/badge/github.com/nar10z/go-collector.svg)](https://pkg.go.dev/github.com/nar10z/go-collector)

Solution for accumulation of events and their subsequent processing.

<img alt="Logo" height="450" src="./image.png" title="Logo"/>

```
go get github.com/nar10z/go-collector
```

## What for?

Sometimes there is a situation where processing data on 1 item is too long.
The [go-collector](https://github.com/nar10z/go-collector) package comes to the rescue!

The solution is to accumulate the data and then process it in a batch. There are 2 situations where the processing
function (**flushFunc**) is called:

- Storage fills up to the maximum value (**flushSize**).
- The interval during which the data is accumulated (**flushInterval**) passes

## Example

```go
package main

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	gocoll "github.com/nar10z/go-collector"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	const (
		countSync  = 4
		countAsync = 3
	)

	collector, err := gocoll.New[string](3, time.Second, func(events []string) error {
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
			errE := collector.AddAsync(ctx, fmt.Sprintf("async №%d", i))
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

				errE := collector.AddSync(ctx, fmt.Sprintf("sync №%d", i))
				if errE != nil {
					fmt.Printf("failed add event: %v\n", errE)
				}
			}()
		}
	}()

	wg.Wait()

	collector.Stop()
}
```

### output:

```text
Start flush 3 events:
 - 1) sync №3
 - 2) sync №0
 - 3) sync №2
Finish
----------------------------------------------------------------------------------------------------
Start flush 3 events:
 - 1) async №0
 - 2) async №1
 - 3) async №2
Finish
----------------------------------------------------------------------------------------------------
Start flush 1 events:
 - 1) sync №1
Finish
----------------------------------------------------------------------------------------------------
```

## License

[MIT](https://raw.githubusercontent.com/nar10z/go-collector/main/LICENSE)
