all: test check bench

test:
	go test -race -v -coverprofile=coverage.txt -covermode=atomic

bench:
	go test -v -run ^Test$$ -bench=. ./... -gocheck.b

check:
	golangci-lint run

.PHONY: test bench check
