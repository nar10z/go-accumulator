# Compare with other packages

### Macbook 13 pro M1
```
goos: darwin
goarch: arm64

Benchmark_accum/#1_go-accumulator,_slice_async-8         	 1000000	       151.1 ns/op	      32 B/op	       2 allocs/op
Benchmark_accum/#2_go-accumulator,_slice_sync-8          	 1000000	       995.1 ns/op	     209 B/op	       5 allocs/op
Benchmark_accum/#3_go-accumulator,_slice-8               	 1000000	       533.6 ns/op	     119 B/op	       3 allocs/op
Benchmark_accum/#4._lrweck/accumulator-8                 	 1000000	       148.2 ns/op	       8 B/op	       1 allocs/op
```


### AMD Ryzen 5 2600 Six-Core Processor
```
goos: windows
goarch: amd64

Benchmark_accum/#1_go-accumulator,_slice_async-12                1000000       1115 ns/op              33 B/op          2 allocs/op
Benchmark_accum/#2_go-accumulator,_slice_sync-12                 1000000       2569 ns/op             209 B/op          5 allocs/op
Benchmark_accum/#3_go-accumulator,_slice-12                      1000000       1827 ns/op             119 B/op          3 allocs/op
Benchmark_accum/#4._lrweck/accumulator-12                        1000000       1589 ns/op               8 B/op          1 allocs/op
```
