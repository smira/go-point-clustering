all: test check

deps:
	go get gopkg.in/check.v1

test:
	go test -v ./... -gocheck.v

bench:
	go test -v -bench=. ./... -gocheck.b

check:
	go tool vet .
	golint .
