.PHONY: test test-cov

test:
	go test ./... -v --race

test-cov:
	go test --tags=tests -coverpkg=./... -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out
	rm coverage.out
