all: test check

deps:
	go get -v -d -t ./...
	go get github.com/golang/lint/golint

test:
	go test -v ./... -gocheck.v

bench:
	go test -v -bench=. ./... -gocheck.b

check:
	go tool vet .
	golint .
