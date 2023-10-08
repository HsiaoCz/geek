package main

import (
	"log"

	"github.com/HsiaoCz/geek/sth/api"
)

const (
	addr = "127.0.0.1:9911"
)

func main() {
	srv := api.NewServer(addr)
	log.Fatal(srv.Start())
}
