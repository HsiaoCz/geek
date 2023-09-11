run: build
	@./bin/img

build:
	@go build -o bin/img main.go

test:
	@go test -v ./...

all: run

help:
	@echo "run: run ./bin/img"
	@echo "build : build porject"
	@echo "test: go test -v ./..."
	@echo "all: just input make"

.PHONY: run build test all help

