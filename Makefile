all: test

test:
	go test -v ./... -gocheck.v

bench:
	go test -v -bench=. ./... -gocheck.b