all: test check

deps:
	go get -v -d -t ./...
	go get github.com/golang/lint/golint
	go get golang.org/x/tools/cmd/vet

test:
	go test -v ./... -gocheck.v

bench:
	go test -v -bench=. ./... -gocheck.b

check:
	go tool vet .
	golint .
