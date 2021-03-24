# Makefile to build the project

COVERAGE = -coverprofile=coverage.txt -covermode=atomic

all: test lint tidy build

travis-ci: test-cov lint tidy

test:
	go test -v -coverpkg=./... -covermode=atomic -coverprofile=cover.out -timeout 2m ./... 2>&1 | tee testoutput.log; go tool cover -func=cover.out

test-cov: 
	go test `go list ./...` ${COVERAGE}

test-int:
	go test `go list ./...` -tags=integration

test-int-cov:
	go test `go list ./...` -tags=integration ${COVERAGE}

lint:
	golangci-lint run

tidy:
	go mod tidy

clean:
	rm -f example

build: example.go
	go build example.go	