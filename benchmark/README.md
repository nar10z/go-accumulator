# Compare with other packages

### Macbook 13 pro M1
```
goos: darwin
goarch: arm64

Benchmark_accum/#1.1_go-accumulator,_list-8         	 1000000	       937.6 ns/op	     160 B/op	       5 allocs/op
Benchmark_accum/#1.2_go-accumulator,_slice-8        	 1000000	       793.8 ns/op	     136 B/op	       4 allocs/op
Benchmark_accum/#1.3_go-accumulator,_stdList-8      	 1000000	       865.4 ns/op	     184 B/op	       5 allocs/op
Benchmark_accum/#2.1_go-accumulator,_list_sync-8    	 1000000	        2019 ns/op	     343 B/op	       8 allocs/op
Benchmark_accum/#2.2_go-accumulator,_slice_sync-8   	 1000000	        1604 ns/op	     312 B/op	       7 allocs/op
Benchmark_accum/#2.3_go-accumulator,_stdList_sync-8 	 1000000	        1792 ns/op	     369 B/op	       8 allocs/op
Benchmark_accum/#3._lrweck/accumulator-8            	 1000000	       147.7 ns/op	       8 B/op	       1 allocs/op
```


### AMD Ryzen 5 2600 Six-Core Processor
```
goos: windows
goarch: amd64

Benchmark_accum
Benchmark_accum/#1.1_go-accumulator,_list-12             1000000              3159 ns/op             160 B/op          5 allocs/op
Benchmark_accum/#1.2_go-accumulator,_slice-12            1000000              2893 ns/op             136 B/op          4 allocs/op
Benchmark_accum/#1.3_go-accumulator,_stdList-12          1000000              3180 ns/op             184 B/op          5 allocs/op
Benchmark_accum/#2.1_go-accumulator,_list_sync-12        1000000              4074 ns/op             349 B/op          8 allocs/op
Benchmark_accum/#2.2_go-accumulator,_slice_sync-12       1000000              4083 ns/op             309 B/op          7 allocs/op
Benchmark_accum/#2.3_go-accumulator,_stdList_sync-12     1000000              4072 ns/op             367 B/op          8 allocs/op
Benchmark_accum/#3._lrweck/accumulator-12                1000000              1533 ns/op               8 B/op          1 allocs/op
```
