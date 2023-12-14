# Compare with other packages

### Macbook 13 pro M1
```
goos: darwin
goarch: arm64

Benchmark_accum/go-accumulator,_async-8         	 1000000	        96.35 ns/op	       8 B/op	       0 allocs/op
Benchmark_accum/lrweck/accumulator-8            	 1000000	       104.3 ns/op	       0 B/op	       0 allocs/op
Benchmark_accum/go-accumulator,_sync-8          	 1000000	       782.7 ns/op	     198 B/op	       4 allocs/op
Benchmark_accum/go-accumulator-8                	 1000000	       423.4 ns/op	     103 B/op	       2 allocs/op
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
