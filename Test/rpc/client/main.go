package main

import (
	"fmt"
	"log"
	"net/rpc"
)

// rpc client
type Args struct {
	X int
	Y int
}

func main() {
	// 建立HTTP连接
	// 基于TCP的rpc调用
	// client, err := rpc.Dial("tcp", "127.0.0.1:9091")
	// if err != nil {
	// 	log.Fatal("dialing:", err)
	// }

	client, err := rpc.DialHTTP("tcp", "127.0.0.1:9091")
	if err != nil {
		log.Fatal(err)
	}

	// 同步调用
	args := &Args{10, 20}
	var reply int
	err = client.Call("ServiceA.Add", args, &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(args.X, args.Y)

	// 异步调用
	var reply2 int
	divCall := client.Go("ServiceA.Add", args, &reply2, nil)
	replyCall := <-divCall.Done // 接收调用结果
	fmt.Println(replyCall.Error)
	fmt.Println(reply2)
}
