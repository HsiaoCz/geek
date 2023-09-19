package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/HsiaoCz/geek/grpc/hello/pb"
	"google.golang.org/grpc"
)

type Hello struct {
	Username string
	Msg      string
	Content  string
}

type Service struct {
	pb.UnimplementedHelloServer
}

func (s *Service) HandleHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	hello := Hello{
		Username: in.Username,
		Msg:      in.Msg,
		Content:  in.Content,
	}
	fmt.Println(hello)
	return &pb.HelloResponse{Username: "lisi", Msg: "hello My man"}, nil
}

const (
	network = "tcp"
	addr    = "127.0.0.1:9991"
)

func main() {
	lis, err := net.Listen(network, addr)
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()                 // 创建grpc服务
	pb.RegisterHelloServer(s, &Service{}) // 在服务端注册服务
	err = s.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}
