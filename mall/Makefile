run: build
	@./bin/mall

build:
	@go build -o bin/mall main.go

test:
	@go test -v ./...

gen:
	@protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative api/v1/mall.proto

all: run

.PHONY: run build test gen all

help:
	@echo "run: run the project"
	@echo "build: build mall"
	@echo "test: go test -v ./..."
	@echo "gen: gen the proto"
