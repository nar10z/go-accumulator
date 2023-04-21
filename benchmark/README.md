# Compare with other packages

```
goos: darwin
goarch: arm64

Benchmark_accum/#1.1_go-accumulator,_channel-8         	 1000000	       172.1 ns/op	      40 B/op	       2 allocs/op
Benchmark_accum/#1.2_go-accumulator,_list-8            	 1000000	       172.4 ns/op	      64 B/op	       3 allocs/op
Benchmark_accum/#1.3_go-accumulator,_slice-8           	 1000000	       187.4 ns/op	      40 B/op	       2 allocs/op
Benchmark_accum/#1.4_go-accumulator,_stdList-8         	 1000000	       156.1 ns/op	      88 B/op	       3 allocs/op
Benchmark_accum/#2._lrweck/accumulator-8               	 1000000	       132.4 ns/op	       8 B/op	       1 allocs/op

```
