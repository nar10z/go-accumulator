# Почему именно каналы?

Потому что они быстрее:

```
BenchmarkStorage
BenchmarkStorage/#1._Channel-8         	        10000000	        65.77 ns/op	      32 B/op	       2 allocs/op
BenchmarkStorage/#2._container/list-8  	        10000000	        74.77 ns/op	      80 B/op	       3 allocs/op
BenchmarkStorage/#3._gods/singlylinkedlist-8    10000000	        75.87 ns/op	      56 B/op	       3 allocs/op
BenchmarkStorage/#4._slice-8                    10000000	        49.72 ns/op	      32 B/op	       2 allocs/op
```
