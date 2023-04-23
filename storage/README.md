# Comparison of different methods of data accumulation

### Macbook 13 pro M1
```
goos: darwin
goarch: arm64

BenchmarkStorage/#1._list-8         	 1000000	       379.9 ns/op	     127 B/op	       4 allocs/op
BenchmarkStorage/#2._gods/list-8    	 1000000	       381.7 ns/op	     103 B/op	       4 allocs/op
BenchmarkStorage/#3._slice-8        	 1000000	       316.2 ns/op	      80 B/op	       3 allocs/op
```

### AMD Ryzen 5 2600 Six-Core Processor
```
goos: windows
goarch: amd64

BenchmarkStorage/#1._list-12             1000000               748.8 ns/op     128 B/op          4 allocs/op
BenchmarkStorage/#2._gods/list-12        1000000               740.2 ns/op     103 B/op          4 allocs/op
BenchmarkStorage/#3._slice-12            1000000               700.3 ns/op      80 B/op          3 allocs/op
PASS```
