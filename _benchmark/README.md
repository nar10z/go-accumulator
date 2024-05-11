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
Benchmark_accum/go-accumulator_async
Benchmark_accum/go-accumulator_async-12                  1000000               153.1 ns/op             0 B/op          0 allocs/op
Benchmark_accum/go-accumulator_async-12                  1000000               132.5 ns/op             0 B/op          0 allocs/op
Benchmark_accum/go-accumulator_async-12                  1000000               197.6 ns/op             0 B/op          0 allocs/op
Benchmark_accum/go-accumulator_async-12                  1000000               128.6 ns/op             0 B/op          0 allocs/op
Benchmark_accum/go-accumulator_async-12                  1000000               177.0 ns/op             0 B/op          0 allocs/op
Benchmark_accum/lrweck/accumulator
Benchmark_accum/lrweck/accumulator-12                    1000000               179.4 ns/op             0 B/op          0 allocs/op
Benchmark_accum/lrweck/accumulator-12                    1000000               174.0 ns/op             0 B/op          0 allocs/op
Benchmark_accum/lrweck/accumulator-12                    1000000               173.9 ns/op             0 B/op          0 allocs/op
Benchmark_accum/lrweck/accumulator-12                    1000000               175.1 ns/op             0 B/op          0 allocs/op
Benchmark_accum/lrweck/accumulator-12                    1000000               169.2 ns/op             0 B/op          0 allocs/op
Benchmark_accum/go-accumulator_sync
Benchmark_accum/go-accumulator_sync-12                   1000000              1760 ns/op             193 B/op          3 allocs/op
Benchmark_accum/go-accumulator_sync-12                   1000000              1736 ns/op             187 B/op          3 allocs/op
Benchmark_accum/go-accumulator_sync-12                   1000000              1705 ns/op             185 B/op          3 allocs/op
Benchmark_accum/go-accumulator_sync-12                   1000000              1766 ns/op             188 B/op          3 allocs/op
Benchmark_accum/go-accumulator_sync-12                   1000000              1790 ns/op             186 B/op          3 allocs/op
Benchmark_accum/go-accumulator
Benchmark_accum/go-accumulator-12                        1000000              1522 ns/op             153 B/op          2 allocs/op
Benchmark_accum/go-accumulator-12                        1000000               678.2 ns/op            56 B/op          0 allocs/op
Benchmark_accum/go-accumulator-12                        1000000               547.7 ns/op            44 B/op          0 allocs/op
Benchmark_accum/go-accumulator-12                        1000000               659.2 ns/op            53 B/op          0 allocs/op
Benchmark_accum/go-accumulator-12                        1000000               694.2 ns/op            62 B/op          1 allocs/op
PASS
```
