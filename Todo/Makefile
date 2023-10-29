run: build
	@./bin/todo

build:
	@go build -o bin/todo main.go

test:
	@go test -v ./...

all: run

.PHONY: run build test all
