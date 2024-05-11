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
