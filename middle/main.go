package main

import "github.com/HsiaoCz/geek/middle/router"

const (
	addr = "127.0.0.1:9091"
)

func main() {
	router.RegisterRouter(addr)
}
