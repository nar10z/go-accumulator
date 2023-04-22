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

Benchmark_accum/#1.1_go-accumulator,_channel-12            1000000       2799 ns/op       40 B/op       2 allocs/op
Benchmark_accum/#1.2_go-accumulator,_list-12               1000000       2718 ns/op       64 B/op       3 allocs/op
Benchmark_accum/#1.3_go-accumulator,_slice-12              1000000       2543 ns/op       40 B/op       2 allocs/op
Benchmark_accum/#1.4_go-accumulator,_stdList-12            1000000       2496 ns/op       90 B/op       3 allocs/op

Benchmark_accum/#2.1_go-accumulator,_channel_sync-12       1000000       3181 ns/op       219 B/op       5 allocs/op
Benchmark_accum/#2.2_go-accumulator,_list_sync-12          1000000       3108 ns/op       241 B/op       6 allocs/op
Benchmark_accum/#2.3_go-accumulator,_slice_sync-12         1000000       3034 ns/op       211 B/op       5 allocs/op
Benchmark_accum/#2.4_go-accumulator,_stdList_sync-12	   1000000       3158 ns/op       265 B/op       6 allocs/op

Benchmark_accum/#3._lrweck/accumulator-12                  1000000       2864 ns/op         8 B/op       1 allocs/op
```
