module accumulator-example

go 1.19

require (
	github.com/lrweck/accumulator v0.0.0-20230204043344-6f6538ed8d35
	github.com/nar10z/go-accumulator v1.0.0
)

require github.com/emirpasic/gods v1.18.1 // indirect

replace (
	github.com/nar10z/go-accumulator => ../
)
