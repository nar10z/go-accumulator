# Compare with other packages

### Macbook 13 pro M1
```
goos: darwin
goarch: arm64

Benchmark_accum/#1.1_go-accumulator,_channel-8       1000000       172.1 ns/op       40 B/op       2 allocs/op
Benchmark_accum/#1.2_go-accumulator,_list-8          1000000       172.4 ns/op       64 B/op       3 allocs/op
Benchmark_accum/#1.3_go-accumulator,_slice-8         1000000	   187.4 ns/op       40 B/op       2 allocs/op
Benchmark_accum/#1.4_go-accumulator,_stdList-8       1000000	   156.1 ns/op       88 B/op       3 allocs/op
Benchmark_accum/#2._lrweck/accumulator-8             1000000	   132.4 ns/op	      8 B/op       1 allocs/op
```

### AMD Ryzen 5 2600 Six-Core Processor
```
goos: windows
goarch: amd64
cpu: AMD Ryzen 5 2600 Six-Core Processor
Benchmark_accum/#1.1_go-accumulator,_channel-12       1000000       2724 ns/op        41 B/op       2 allocs/op
Benchmark_accum/#1.2_go-accumulator,_list-12          1000000       2514 ns/op        66 B/op       3 allocs/op
Benchmark_accum/#1.3_go-accumulator,_slice-12         1000000       1296 ns/op        41 B/op       2 allocs/op
Benchmark_accum/#1.4_go-accumulator,_stdList-12       1000000       2261 ns/op        90 B/op       3 allocs/op
Benchmark_accum/#1.5_go-accumulator,_stdList_sync-12  1000000       3189 ns/op       269 B/op       6 allocs/op
Benchmark_accum/#2._lrweck/accumulator-12             1000000       2854 ns/op         8 B/op       1 allocs/op
```
