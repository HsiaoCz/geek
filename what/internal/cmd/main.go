package main

import (
	"log"

	"github.com/HsiaoCz/geek/what/internal/server"
)

var (
	network = "tcp"
	addr    = "127.0.0.1:3002"
)

func main() {
	log.Println("The server is running on port:", addr)
	if err := server.RegGrpc(network, addr); err != nil {
		log.Fatal(err)
	}
}
