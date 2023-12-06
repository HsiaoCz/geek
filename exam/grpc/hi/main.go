package main

import (
	"context"
	"log"
	"time"

	"github.com/HsiaoCz/geek/exam/grpc/hi/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = "127.0.0.1:3001"
)

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	c := pb.NewHelloClient(conn)

	// 执行调用并打印响应数据
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.HandleHello(ctx, &pb.HelloRequest{Username: "zhangsan", Msg: "hello my man", Content: "what's up"})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(r.GetUsername(), r.GetMsg())
}
