run: build
	@./bin/tenz

build: 
	@go build -o bin/tenz main.go

test:
	@go test -v ./...

all: run

.PHONY: run build test all
