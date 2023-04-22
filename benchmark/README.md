# Compare with other packages

### Macbook 13 pro M1
```
goos: darwin
goarch: arm64

Benchmark_accum/#1.1_go-accumulator,_channel-8         	 1000000	       195.6 ns/op	      90 B/op	       3 allocs/op
Benchmark_accum/#1.2_go-accumulator,_list-8            	 1000000	       182.2 ns/op	      66 B/op	       3 allocs/op
Benchmark_accum/#1.3_go-accumulator,_slice-8           	 1000000	       203.4 ns/op	      75 B/op	       2 allocs/op
Benchmark_accum/#1.4_go-accumulator,_stdList-8         	 1000000	       206.2 ns/op	      90 B/op	       3 allocs/op

Benchmark_accum/#2.1_go-accumulator,_channel_sync-8    	 1000000	        1589 ns/op	     271 B/op	       6 allocs/op
Benchmark_accum/#2.2_go-accumulator,_list_sync-8       	 1000000	        1226 ns/op	     239 B/op	       6 allocs/op
Benchmark_accum/#2.3_go-accumulator,_slice_sync-8      	 1000000	        1199 ns/op	     247 B/op	       5 allocs/op
Benchmark_accum/#2.4_go-accumulator,_stdList_sync-8    	 1000000	        1255 ns/op	     266 B/op	       6 allocs/op

Benchmark_accum/#3._lrweck/accumulator-8               	 1000000	       138.4 ns/op	       8 B/op	       1 allocs/op
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
