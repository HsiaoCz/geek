package main

import "github.com/HsiaoCz/geek/img/router"

const (
	addr = "127.0.0.1:9991"
)

func main() {
	router.RegRouter(addr)
}
