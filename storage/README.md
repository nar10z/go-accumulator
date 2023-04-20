# Storage


If speed of **execution** is important to you, but _RAM_ is not so important - it is better to use a storage that uses a channel.

If you have limited **RAM**, you can use a repository that uses a single linked list.

## Comparison of different methods of data accumulation

### Macbook 13 pro M1
```
goos: darwin
goarch: arm64

BenchmarkStorage/#1._Channel-8         	        10000000	        65.77 ns/op	      32 B/op	       2 allocs/op
BenchmarkStorage/#2._container/list-8  	        10000000	        74.77 ns/op	      80 B/op	       3 allocs/op
BenchmarkStorage/#3._gods/singlylinkedlist-8    10000000	        75.87 ns/op	      56 B/op	       3 allocs/op
BenchmarkStorage/#4._slice-8                    10000000	        49.72 ns/op	      32 B/op	       2 allocs/op
```

### AMD Ryzen 5 2600 Six-Core Processor
```
goos: windows
goarch: amd64

BenchmarkStorage/#1._Channel-12                 10000000               147.5 ns/op            32 B/op          2 allocs/op
BenchmarkStorage/#2._container/list-12          10000000               170.6 ns/op            81 B/op          3 allocs/op
BenchmarkStorage/#3._gods/singlylinkedlist-12   10000000               173.4 ns/op            57 B/op          3 allocs/op
BenchmarkStorage/#4._slice-12                   10000000               106.2 ns/op            32 B/op          2 allocs/op
```
