package main

import (
	"fmt"

	"github.com/HsiaoCz/geek/Test/structure/api"
)

const (
	addr = "127.0.0.1:3033"
)

func main() {
	fmt.Println("hello mono")
	s := api.NewAPIServer(addr)
	s.Run()
}
