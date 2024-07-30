# Compare with other packages

```
go test -bench=. -test.benchmem -test.benchtime=1000000x -test.count=5
```

### Macbook 13 pro M1
```
goos: darwin
goarch: arm64
pkg: accumulator-example
Benchmark_accum
Benchmark_accum/go-accumulator_async
Benchmark_accum/go-accumulator_async-8         	 1000000	       167.0 ns/op	       0 B/op	       0 allocs/op
Benchmark_accum/go-accumulator_async-8         	 1000000	       167.8 ns/op	       0 B/op	       0 allocs/op
Benchmark_accum/go-accumulator_async-8         	 1000000	       189.2 ns/op	       0 B/op	       0 allocs/op
Benchmark_accum/go-accumulator_async-8         	 1000000	       105.1 ns/op	       0 B/op	       0 allocs/op
Benchmark_accum/go-accumulator_async-8         	 1000000	       104.8 ns/op	       0 B/op	       0 allocs/op
Benchmark_accum/lrweck/accumulator
Benchmark_accum/lrweck/accumulator-8           	 1000000	       146.5 ns/op	       0 B/op	       0 allocs/op
Benchmark_accum/lrweck/accumulator-8           	 1000000	       143.3 ns/op	       0 B/op	       0 allocs/op
Benchmark_accum/lrweck/accumulator-8           	 1000000	       142.8 ns/op	       0 B/op	       0 allocs/op
Benchmark_accum/lrweck/accumulator-8           	 1000000	       140.9 ns/op	       0 B/op	       0 allocs/op
Benchmark_accum/lrweck/accumulator-8           	 1000000	       144.5 ns/op	       0 B/op	       0 allocs/op
Benchmark_accum/go-accumulator_sync
Benchmark_accum/go-accumulator_sync-8          	 1000000	      1352 ns/op	     187 B/op	       3 allocs/op
Benchmark_accum/go-accumulator_sync-8          	 1000000	      1353 ns/op	     182 B/op	       3 allocs/op
Benchmark_accum/go-accumulator_sync-8          	 1000000	      1303 ns/op	     184 B/op	       3 allocs/op
Benchmark_accum/go-accumulator_sync-8          	 1000000	      1302 ns/op	     182 B/op	       3 allocs/op
Benchmark_accum/go-accumulator_sync-8          	 1000000	      1303 ns/op	     183 B/op	       3 allocs/op
Benchmark_accum/go-accumulator
Benchmark_accum/go-accumulator-8               	 1000000	       882.8 ns/op	      89 B/op	       1 allocs/op
Benchmark_accum/go-accumulator-8               	 1000000	       381.6 ns/op	      26 B/op	       0 allocs/op
Benchmark_accum/go-accumulator-8               	 1000000	      1403 ns/op	     179 B/op	       3 allocs/op
Benchmark_accum/go-accumulator-8               	 1000000	      1103 ns/op	     141 B/op	       2 allocs/op
Benchmark_accum/go-accumulator-8               	 1000000	       233.2 ns/op	       4 B/op	       0 allocs/op
PASS
```


### AMD Ryzen 5 2600 Six-Core Processor
```
goos: windows
goarch: amd64
pkg: accumulator-example
cpu: AMD Ryzen 5 2600 Six-Core Processor
Benchmark_accum
Benchmark_accum/go-accumulator_async
Benchmark_accum/go-accumulator_async-12         	  100000	       176.7 ns/op	       4 B/op	       0 allocs/op
Benchmark_accum/go-accumulator_async-12         	  100000	       166.6 ns/op	       7 B/op	       0 allocs/op
Benchmark_accum/go-accumulator_async-12         	  100000	       140.0 ns/op	       4 B/op	       0 allocs/op
Benchmark_accum/go-accumulator_async-12         	  100000	       140.1 ns/op	       4 B/op	       0 allocs/op
Benchmark_accum/go-accumulator_async-12         	  100000	       180.0 ns/op	       5 B/op	       0 allocs/op
Benchmark_accum/lrweck/accumulator
Benchmark_accum/lrweck/accumulator-12           	  100000	       170.1 ns/op	       1 B/op	       0 allocs/op
Benchmark_accum/lrweck/accumulator-12           	  100000	       169.9 ns/op	       1 B/op	       0 allocs/op
Benchmark_accum/lrweck/accumulator-12           	  100000	       170.1 ns/op	       1 B/op	       0 allocs/op
Benchmark_accum/lrweck/accumulator-12           	  100000	       169.9 ns/op	       1 B/op	       0 allocs/op
Benchmark_accum/lrweck/accumulator-12           	  100000	       190.0 ns/op	       1 B/op	       0 allocs/op
Benchmark_accum/go-accumulator_sync
Benchmark_accum/go-accumulator_sync-12          	  100000	      5030 ns/op	     240 B/op	       3 allocs/op
Benchmark_accum/go-accumulator_sync-12          	  100000	      2530 ns/op	     209 B/op	       3 allocs/op
Benchmark_accum/go-accumulator_sync-12          	  100000	      2520 ns/op	     202 B/op	       3 allocs/op
Benchmark_accum/go-accumulator_sync-12          	  100000	      2520 ns/op	     207 B/op	       3 allocs/op
Benchmark_accum/go-accumulator_sync-12          	  100000	      2520 ns/op	     206 B/op	       3 allocs/op
Benchmark_accum/go-accumulator
Benchmark_accum/go-accumulator-12               	  100000	      1510 ns/op	      94 B/op	       1 allocs/op
Benchmark_accum/go-accumulator-12               	  100000	      1001 ns/op	      89 B/op	       1 allocs/op
Benchmark_accum/go-accumulator-12               	  100000	       480.0 ns/op	      50 B/op	       0 allocs/op
Benchmark_accum/go-accumulator-12               	  100000	       149.9 ns/op	       7 B/op	       0 allocs/op
Benchmark_accum/go-accumulator-12               	  100000	      2530 ns/op	     205 B/op	       3 allocs/op
PASS
```
