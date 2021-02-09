all: build test lint coverage

build:
	go build ./...

test:
	go test ./...

lint:
	golangci-lint run

coverage:
	go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out