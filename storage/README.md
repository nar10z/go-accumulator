# Comparison of different methods of data accumulation

### Macbook 13 pro M1
```
goos: darwin
goarch: arm64

BenchmarkStorage/#1._Channel-8         	        10000000	        65.77 ns/op	      32 B/op	       2 allocs/op
BenchmarkStorage/#2._container/list-8      	    10000000	        74.77 ns/op	      80 B/op	       3 allocs/op
BenchmarkStorage/#3._gods/singlylinkedlist-8    10000000	        75.87 ns/op	      56 B/op	       3 allocs/op
BenchmarkStorage/#4._slice-8                    10000000	        49.72 ns/op	      32 B/op	       2 allocs/op
```

### AMD Ryzen 5 2600 Six-Core Processor
```
goos: windows
goarch: amd64

BenchmarkStorage/#1._list-12             1000000      158.1 ns/op      79 B/op      2 allocs/op
BenchmarkStorage/#2._gods/list-12        1000000      166.7 ns/op      55 B/op      2 allocs/op
BenchmarkStorage/#3._slice-12            1000000      112.9 ns/op      32 B/op      2 allocs/op
```
