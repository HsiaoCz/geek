run: build
	@./bin/middle

build:
	@go build -o bin/middle main.go

test:
	@go test -v ./...

.PHONY: run test build

all: run

help:
	@echo "run: run the project"
	@echo "build: go build"
	@echo "test: go test "