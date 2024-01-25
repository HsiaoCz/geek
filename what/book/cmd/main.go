package main

import (
	"log"

	"github.com/HsiaoCz/geek/what/book/server"
)

var (
	network = "tcp"
	adr     = "127.0.0.1:3003"
)

func main() {
	if err := server.RegisterGrpc(adr, network); err != nil {
		log.Fatal(err)
	}
}
