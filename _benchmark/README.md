# Compare with other packages

```
go test -bench=. -test.benchmem -test.benchtime=1000000x -test.count=5
```

### Macbook 13 pro M1
```
goos: darwin
goarch: arm64

Benchmark_accum/go-accumulator,_async
Benchmark_accum/go-accumulator,_async-8         	 1000000	       159.7 ns/op	       0 B/op	       0 allocs/op
Benchmark_accum/go-accumulator,_async-8         	 1000000	       152.2 ns/op	       0 B/op	       0 allocs/op
Benchmark_accum/go-accumulator,_async-8         	 1000000	       153.4 ns/op	       0 B/op	       0 allocs/op
Benchmark_accum/go-accumulator,_async-8         	 1000000	       153.5 ns/op	       0 B/op	       0 allocs/op
Benchmark_accum/go-accumulator,_async-8         	 1000000	        85.99 ns/op	       0 B/op	       0 allocs/op
Benchmark_accum/lrweck/accumulator
Benchmark_accum/lrweck/accumulator-8            	 1000000	       141.6 ns/op	       0 B/op	       0 allocs/op
Benchmark_accum/lrweck/accumulator-8            	 1000000	       142.4 ns/op	       0 B/op	       0 allocs/op
Benchmark_accum/lrweck/accumulator-8            	 1000000	       144.2 ns/op	       0 B/op	       0 allocs/op
Benchmark_accum/lrweck/accumulator-8            	 1000000	       130.5 ns/op	       0 B/op	       0 allocs/op
Benchmark_accum/lrweck/accumulator-8            	 1000000	       154.5 ns/op	       0 B/op	       0 allocs/op
Benchmark_accum/go-accumulator,_sync
Benchmark_accum/go-accumulator,_sync-8          	 1000000	       982.0 ns/op	     195 B/op	       3 allocs/op
Benchmark_accum/go-accumulator,_sync-8          	 1000000	       896.2 ns/op	     186 B/op	       3 allocs/op
Benchmark_accum/go-accumulator,_sync-8          	 1000000	       899.6 ns/op	     188 B/op	       3 allocs/op
Benchmark_accum/go-accumulator,_sync-8          	 1000000	       848.5 ns/op	     188 B/op	       3 allocs/op
Benchmark_accum/go-accumulator,_sync-8          	 1000000	       838.5 ns/op	     185 B/op	       3 allocs/op
Benchmark_accum/go-accumulator
Benchmark_accum/go-accumulator-8                	 1000000	       426.2 ns/op	      64 B/op	       1 allocs/op
Benchmark_accum/go-accumulator-8                	 1000000	       617.4 ns/op	     116 B/op	       1 allocs/op
Benchmark_accum/go-accumulator-8                	 1000000	       611.4 ns/op	     120 B/op	       2 allocs/op
Benchmark_accum/go-accumulator-8                	 1000000	       310.1 ns/op	      36 B/op	       0 allocs/op
Benchmark_accum/go-accumulator-8                	 1000000	       195.0 ns/op	      16 B/op	       0 allocs/op
PASS
```


### AMD Ryzen 5 2600 Six-Core Processor
```
goos: windows
goarch: amd64

Benchmark_accum/go-accumulator,_async-12                 1000000               228.6 ns/op             8 B/op          0 allocs/op
Benchmark_accum/lrweck/accumulator-12                    1000000               217.6 ns/op             0 B/op          0 allocs/op
Benchmark_accum/go-accumulator,_sync-12                  1000000              1482 ns/op             179 B/op          3 allocs/op
Benchmark_accum/go-accumulator-12                        1000000               883.9 ns/op            93 B/op          1 allocs/op
```
