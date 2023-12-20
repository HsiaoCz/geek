package main

import (
	"context"
	"fmt"
	"log"

	"github.com/HsiaoCz/geek/exam/grpc/mate-c/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func UserLogin(c pb.GreeterClient) {
	md := metadata.Pairs(
		"token", "app-test-hsiaocz",
		"request_id", "123456",
	)
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	res, err := c.UserLogin(ctx, &pb.LoginRequest{Username: "zhangsan", Passwrod: "12334", Repassword: "12334", Email: "12334445@qq.com"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}

var (
	addr = "127.0.0.1:9001"
)

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	cc := pb.NewGreeterClient(conn)
	UserLogin(cc)
}
