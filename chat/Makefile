run: build
	@./bin/chat

build:
	@go build -o bin/chat main.go

test:
	@go test -v ./...

all: run

.PHONY: run build test all

help:
	@echo "run: run the project"
	@echo "build: go build -o bin/chat"
	@echo "test: go test -v ./..."
	@echo "all: enter this will use run"
