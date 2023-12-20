package main

import (
	"log"

	"github.com/HsiaoCz/geek/what/route"
)

var (
	addr = "127.0.0.1:3001"
)

func main() {
	if err := route.Router(addr); err != nil {
		log.Fatal(err)
	}
}
