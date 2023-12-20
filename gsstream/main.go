package main

import (
	"fmt"
	"log"
)

// underlying storage (in memory)
// server (http,tcp)
func main() {
	cfg := &Config{
		ListenAddr: "127.0.0.1:9001",
	}
	s, err := NewServer(cfg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(s)
	fmt.Println("everything is ok")
}
