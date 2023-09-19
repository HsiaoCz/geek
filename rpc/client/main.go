package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Args struct {
	X int
	Y int
}

const (
	network = "tcp"
	addr    = "127.0.0.1:9991"
)

func main() {
	client, err := rpc.DialHTTP(network, addr)
	if err != nil {
		log.Fatal(err)
	}

	// 同步调用
	args := &Args{10, 20}
	var reply int
	err = client.Call("Service.Add", args, &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(args.X, args.Y)
}
