# Почему именно каналы?

Потому что они быстрее:

```
goos: darwin
goarch: arm64
BenchmarkStorage
BenchmarkStorage/#1._Channel-8         	        10000000	        65.77 ns/op	      32 B/op	       2 allocs/op
BenchmarkStorage/#2._container/list-8  	        10000000	        74.77 ns/op	      80 B/op	       3 allocs/op
BenchmarkStorage/#3._gods/singlylinkedlist-8    10000000	        75.87 ns/op	      56 B/op	       3 allocs/op
BenchmarkStorage/#4._slice-8                    10000000	        49.72 ns/op	      32 B/op	       2 allocs/op
```

```
goos: windows
goarch: amd64
cpu: AMD Ryzen 5 2600 Six-Core Processor
BenchmarkStorage/#1._Channel-12                 10000000               146.8 ns/op
BenchmarkStorage/#2._container/list-12          10000000               167.5 ns/op
BenchmarkStorage/#3._gods/singlylinkedlist-12   10000000               172.3 ns/op
BenchmarkStorage/#4._slice-12                   10000000               106.9 ns/op
PASS
```
