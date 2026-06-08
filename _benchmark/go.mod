module accumulator-example

go 1.26

require (
	github.com/lrweck/accumulator v0.0.0-20230204043344-6f6538ed8d35
	github.com/nar10z/go-accumulator/v2 v2.0.0
	golang.org/x/sync v0.20.0
)

require (
	github.com/bytedance/gopkg v0.1.4 // indirect
	golang.org/x/sys v0.30.0 // indirect
)

replace github.com/nar10z/go-accumulator => ../
