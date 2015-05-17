ifeq ($(GOVERSION), devel)
TRAVIS_TARGET=coveralls
else
TRAVIS_TARGET=test
endif

all: test check

deps:
	go get -v -d -t ./...
	go get github.com/golang/lint/golint
	go get golang.org/x/tools/cmd/vet
	go get github.com/mattn/goveralls

test:
	go test -v ./... -gocheck.v

bench:
	go test -v -bench=. ./... -gocheck.b

coverage.out:
	go test -coverprofile=coverage.out -covermode=count .

coverage: coverage.out
	go tool cover -html=coverage.out
	rm -f coverage.out

coveralls: coverage.out
	goveralls -service coveralls -coverprofile=coverage.out $(COVERALLS_TOKEN)
	rm -f coverage.out

check:
	go tool vet .
	golint .

travis: $(TRAVIS_TARGET)