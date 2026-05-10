# Compare with other packages

```
go test -bench=. -test.benchmem -test.benchtime=1000000x -test.count=5
```

### Macbook 13 pro M1
```
goos: darwin
goarch: arm64
pkg: accumulator-example
cpu: Apple M1
Benchmark_accum
Benchmark_accum/go-accumulator_async
Benchmark_accum/go-accumulator_async-8         	 1000000	       152.6 ns/op	       0 B/op	       0 allocs/op
Benchmark_accum/go-accumulator_async-8         	 1000000	       155.5 ns/op	       0 B/op	       0 allocs/op
Benchmark_accum/go-accumulator_async-8         	 1000000	       242.0 ns/op	       0 B/op	       0 allocs/op
Benchmark_accum/go-accumulator_async-8         	 1000000	       187.9 ns/op	       0 B/op	       0 allocs/op
Benchmark_accum/go-accumulator_async-8         	 1000000	       199.1 ns/op	       0 B/op	       0 allocs/op
Benchmark_accum/lrweck/accumulator
Benchmark_accum/lrweck/accumulator-8           	 1000000	       149.0 ns/op	       0 B/op	       0 allocs/op
Benchmark_accum/lrweck/accumulator-8           	 1000000	       151.2 ns/op	       0 B/op	       0 allocs/op
Benchmark_accum/lrweck/accumulator-8           	 1000000	       148.7 ns/op	       0 B/op	       0 allocs/op
Benchmark_accum/lrweck/accumulator-8           	 1000000	       158.8 ns/op	       0 B/op	       0 allocs/op
Benchmark_accum/lrweck/accumulator-8           	 1000000	       152.6 ns/op	       0 B/op	       0 allocs/op
Benchmark_accum/go-accumulator_sync
Benchmark_accum/go-accumulator_sync-8          	 1000000	      1053 ns/op	     209 B/op	       4 allocs/op
Benchmark_accum/go-accumulator_sync-8          	 1000000	      1052 ns/op	     206 B/op	       4 allocs/op
Benchmark_accum/go-accumulator_sync-8          	 1000000	      1102 ns/op	     204 B/op	       4 allocs/op
Benchmark_accum/go-accumulator_sync-8          	 1000000	      1101 ns/op	     203 B/op	       4 allocs/op
Benchmark_accum/go-accumulator_sync-8          	 1000000	      1052 ns/op	     205 B/op	       4 allocs/op
Benchmark_accum/go-accumulator
Benchmark_accum/go-accumulator-8               	 1000000	       462.2 ns/op	      82 B/op	       1 allocs/op
Benchmark_accum/go-accumulator-8               	 1000000	       901.7 ns/op	     175 B/op	       3 allocs/op
Benchmark_accum/go-accumulator-8               	 1000000	       851.7 ns/op	     155 B/op	       3 allocs/op
Benchmark_accum/go-accumulator-8               	 1000000	       353.9 ns/op	      55 B/op	       1 allocs/op
Benchmark_accum/go-accumulator-8               	 1000000	       205.3 ns/op	       8 B/op	       0 allocs/op
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
