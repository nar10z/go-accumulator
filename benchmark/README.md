# Compare with other packages

### Macbook 13 pro M1
```
goos: darwin
goarch: arm64

Benchmark_accum/#1.1_go-accumulator,_list-8         	 1000000	       296.8 ns/op	      64 B/op	       3 allocs/op
Benchmark_accum/#1.2_go-accumulator,_slice-8        	 1000000	       253.8 ns/op	      40 B/op	       2 allocs/op
Benchmark_accum/#1.3_go-accumulator,_stdList-8      	 1000000	       297.8 ns/op	      88 B/op	       3 allocs/op

Benchmark_accum/#2.1_go-accumulator,_list_sync-8    	 1000000	      2015 ns/op	     248 B/op	       6 allocs/op
Benchmark_accum/#2.2_go-accumulator,_slice_sync-8   	 1000000	      2017 ns/op	     219 B/op	       5 allocs/op
Benchmark_accum/#2.3_go-accumulator,_stdList_sync-8 	 1000000	      2016 ns/op	     261 B/op	       6 allocs/op

Benchmark_accum/#3._lrweck/accumulator-8            	 1000000	       181.2 ns/op	       8 B/op	       1 allocs/op
```


### AMD Ryzen 5 2600 Six-Core Processor
```
goos: windows
goarch: amd64

Benchmark_accum/#1.1_go-accumulator,_list-12             1000000              1590 ns/op              64 B/op          3 allocs/op
Benchmark_accum/#1.2_go-accumulator,_slice-12            1000000              1570 ns/op              40 B/op          2 allocs/op
Benchmark_accum/#1.3_go-accumulator,_stdList-12          1000000              1600 ns/op              88 B/op          3 allocs/op

Benchmark_accum/#2.1_go-accumulator,_list_sync-12        1000000              3041 ns/op             249 B/op          6 allocs/op
Benchmark_accum/#2.2_go-accumulator,_slice_sync-12       1000000              3050 ns/op             223 B/op          5 allocs/op
Benchmark_accum/#2.3_go-accumulator,_stdList_sync-12     1000000              3034 ns/op             267 B/op          6 allocs/op

Benchmark_accum/#3._lrweck/accumulator-12                1000000              1538 ns/op               8 B/op          1 allocs/op
```
