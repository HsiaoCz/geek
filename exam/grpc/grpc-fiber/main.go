package main

import (
	"log"

	"github.com/HsiaoCz/geek/exam/grpc/grpc-fiber/server"
)

var (
	net_work   = "tcp"
	fiber_addr = "127.0.0.1:3001"
	grpc_addr  = "127.0.0.1:3003"
)

func main() {
	if err := server.RegisterHttpServer(fiber_addr); err != nil {
		log.Fatal(err)
	}
	if err := server.RegisterGrpcServer(net_work, grpc_addr); err != nil {
		log.Fatal(err)
	}
}
