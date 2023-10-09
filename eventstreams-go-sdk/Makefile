# Makefile to build the project

COVERAGE = -coverprofile=coverage.txt -covermode=atomic
ADMINREST_EXAMPLE_DIR = examples/adminrestv1
SCHEMA_EXAMPLE_DIR = examples/schemaregistryv1

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
	rm -f ${ADMINREST_EXAMPLE_DIR}/example
	rm -f ${SCHEMA_EXAMPLE_DIR}/example

adminrest-build: ${ADMINREST_EXAMPLE_DIR}/example.go
	cd ${ADMINREST_EXAMPLE_DIR} && go build -o example	

schema-build: ${SCHEMA_EXAMPLE_DIR}/example.go
	cd ${SCHEMA_EXAMPLE_DIR} && go build -o example

build: adminrest-build schema-build