package main

import "fmt"

const (
	addr = "127.0.0.1:3033"
)

func main() {
	fmt.Println("hello mono")
	s := NewAPIServer(addr)
	s.Run()
}
