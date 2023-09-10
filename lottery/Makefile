run: build
	@./bin/lottery

build:
	@go build -o bin/lottery main.go

test:
	@go test -v ./...

all: run

help:
	@echo "run: run lottery"
	@echo "build: go build -o lottery"
	@echo "test: go test -v ./..."

.PHONY: run build test all
