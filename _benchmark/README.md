# Compare with other packages

```
go test -bench=. -test.benchmem -test.benchtime=1000000x -test.count=5
```

### Macbook 13 pro M1
```
goos: darwin
goarch: arm64

Benchmark_accum/go-accumulator_async
Benchmark_accum/go-accumulator_async-8         	 1000000	       110.1 ns/op	       0 B/op	       0 allocs/op
Benchmark_accum/go-accumulator_async-8         	 1000000	       176.3 ns/op	       0 B/op	       0 allocs/op
Benchmark_accum/go-accumulator_async-8         	 1000000	       161.3 ns/op	       0 B/op	       0 allocs/op
Benchmark_accum/go-accumulator_async-8         	 1000000	       130.6 ns/op	       0 B/op	       0 allocs/op
Benchmark_accum/go-accumulator_async-8         	 1000000	       123.9 ns/op	       0 B/op	       0 allocs/op
Benchmark_accum/lrweck/accumulator
Benchmark_accum/lrweck/accumulator-8           	 1000000	       183.8 ns/op	       0 B/op	       0 allocs/op
Benchmark_accum/lrweck/accumulator-8           	 1000000	       188.3 ns/op	       0 B/op	       0 allocs/op
Benchmark_accum/lrweck/accumulator-8           	 1000000	       185.2 ns/op	       0 B/op	       0 allocs/op
Benchmark_accum/lrweck/accumulator-8           	 1000000	       187.2 ns/op	       0 B/op	       0 allocs/op
Benchmark_accum/lrweck/accumulator-8           	 1000000	       187.4 ns/op	       0 B/op	       0 allocs/op
Benchmark_accum/go-accumulator_sync
Benchmark_accum/go-accumulator_sync-8          	 1000000	       883.1 ns/op	     188 B/op	       3 allocs/op
Benchmark_accum/go-accumulator_sync-8          	 1000000	       829.5 ns/op	     186 B/op	       3 allocs/op
Benchmark_accum/go-accumulator_sync-8          	 1000000	       905.8 ns/op	     187 B/op	       3 allocs/op
Benchmark_accum/go-accumulator_sync-8          	 1000000	       857.7 ns/op	     182 B/op	       3 allocs/op
Benchmark_accum/go-accumulator_sync-8          	 1000000	       922.9 ns/op	     186 B/op	       3 allocs/op
Benchmark_accum/go-accumulator
Benchmark_accum/go-accumulator-8               	 1000000	       183.1 ns/op	       1 B/op	       0 allocs/op
Benchmark_accum/go-accumulator-8               	 1000000	       673.8 ns/op	     136 B/op	       2 allocs/op
Benchmark_accum/go-accumulator-8               	 1000000	       569.5 ns/op	     107 B/op	       1 allocs/op
Benchmark_accum/go-accumulator-8               	 1000000	       269.8 ns/op	      30 B/op	       0 allocs/op
Benchmark_accum/go-accumulator-8               	 1000000	       715.6 ns/op	     105 B/op	       1 allocs/op
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
