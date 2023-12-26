# Compare with other packages

```
go test -bench=. -test.benchmem -test.benchtime=1000000x -test.count=5
```

### Macbook 13 pro M1
```
goos: darwin
goarch: arm64

Benchmark_accum/go-accumulator,_async
Benchmark_accum/go-accumulator,_async-8         	 1000000	       159.7 ns/op	       0 B/op	       0 allocs/op
Benchmark_accum/go-accumulator,_async-8         	 1000000	       152.2 ns/op	       0 B/op	       0 allocs/op
Benchmark_accum/go-accumulator,_async-8         	 1000000	       153.4 ns/op	       0 B/op	       0 allocs/op
Benchmark_accum/go-accumulator,_async-8         	 1000000	       153.5 ns/op	       0 B/op	       0 allocs/op
Benchmark_accum/go-accumulator,_async-8         	 1000000	        85.99 ns/op	       0 B/op	       0 allocs/op
Benchmark_accum/lrweck/accumulator
Benchmark_accum/lrweck/accumulator-8            	 1000000	       141.6 ns/op	       0 B/op	       0 allocs/op
Benchmark_accum/lrweck/accumulator-8            	 1000000	       142.4 ns/op	       0 B/op	       0 allocs/op
Benchmark_accum/lrweck/accumulator-8            	 1000000	       144.2 ns/op	       0 B/op	       0 allocs/op
Benchmark_accum/lrweck/accumulator-8            	 1000000	       130.5 ns/op	       0 B/op	       0 allocs/op
Benchmark_accum/lrweck/accumulator-8            	 1000000	       154.5 ns/op	       0 B/op	       0 allocs/op
Benchmark_accum/go-accumulator,_sync
Benchmark_accum/go-accumulator,_sync-8          	 1000000	       982.0 ns/op	     195 B/op	       3 allocs/op
Benchmark_accum/go-accumulator,_sync-8          	 1000000	       896.2 ns/op	     186 B/op	       3 allocs/op
Benchmark_accum/go-accumulator,_sync-8          	 1000000	       899.6 ns/op	     188 B/op	       3 allocs/op
Benchmark_accum/go-accumulator,_sync-8          	 1000000	       848.5 ns/op	     188 B/op	       3 allocs/op
Benchmark_accum/go-accumulator,_sync-8          	 1000000	       838.5 ns/op	     185 B/op	       3 allocs/op
Benchmark_accum/go-accumulator
Benchmark_accum/go-accumulator-8                	 1000000	       426.2 ns/op	      64 B/op	       1 allocs/op
Benchmark_accum/go-accumulator-8                	 1000000	       617.4 ns/op	     116 B/op	       1 allocs/op
Benchmark_accum/go-accumulator-8                	 1000000	       611.4 ns/op	     120 B/op	       2 allocs/op
Benchmark_accum/go-accumulator-8                	 1000000	       310.1 ns/op	      36 B/op	       0 allocs/op
Benchmark_accum/go-accumulator-8                	 1000000	       195.0 ns/op	      16 B/op	       0 allocs/op
PASS
```


### AMD Ryzen 5 2600 Six-Core Processor
```
goos: windows
goarch: amd64

Benchmark_accum
Benchmark_accum/go-accumulator,_async
Benchmark_accum/go-accumulator,_async-12                 1000000               152.1 ns/op             0 B/op          0 allocs/op
Benchmark_accum/go-accumulator,_async-12                 1000000               197.0 ns/op             0 B/op          0 allocs/op
Benchmark_accum/go-accumulator,_async-12                 1000000               133.0 ns/op             0 B/op          0 allocs/op
Benchmark_accum/go-accumulator,_async-12                 1000000               151.5 ns/op             0 B/op          0 allocs/op
Benchmark_accum/go-accumulator,_async-12                 1000000               177.0 ns/op             0 B/op          0 allocs/op
Benchmark_accum/lrweck/accumulator
Benchmark_accum/lrweck/accumulator-12                    1000000               172.0 ns/op             0 B/op          0 allocs/op
Benchmark_accum/lrweck/accumulator-12                    1000000               163.0 ns/op             0 B/op          0 allocs/op
Benchmark_accum/lrweck/accumulator-12                    1000000               163.0 ns/op             0 B/op          0 allocs/op
Benchmark_accum/lrweck/accumulator-12                    1000000               162.0 ns/op             0 B/op          0 allocs/op
Benchmark_accum/lrweck/accumulator-12                    1000000               170.0 ns/op             0 B/op          0 allocs/op
Benchmark_accum/go-accumulator,_sync
Benchmark_accum/go-accumulator,_sync-12                  1000000              1819 ns/op             192 B/op          3 allocs/op
Benchmark_accum/go-accumulator,_sync-12                  1000000              1807 ns/op             190 B/op          3 allocs/op
Benchmark_accum/go-accumulator,_sync-12                  1000000              1661 ns/op             184 B/op          3 allocs/op
Benchmark_accum/go-accumulator,_sync-12                  1000000              1662 ns/op             185 B/op          3 allocs/op
Benchmark_accum/go-accumulator,_sync-12                  1000000              1734 ns/op             186 B/op          3 allocs/op
Benchmark_accum/go-accumulator
Benchmark_accum/go-accumulator-12                        1000000              1175 ns/op             112 B/op          1 allocs/op
Benchmark_accum/go-accumulator-12                        1000000               718.0 ns/op            63 B/op          1 allocs/op
Benchmark_accum/go-accumulator-12                        1000000               313.0 ns/op            19 B/op          0 allocs/op
Benchmark_accum/go-accumulator-12                        1000000               914.1 ns/op            82 B/op          1 allocs/op
Benchmark_accum/go-accumulator-12                        1000000              1358 ns/op             125 B/op          2 allocs/op
PASS
```
