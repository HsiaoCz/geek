package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/HsiaoCz/geek/exam/grpc/hello/pb"
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
		Username: in.GetUsername(),
		Msg:      in.GetMsg(),
		Content:  in.GetContent(),
	}
	fmt.Println(hello)
	return &pb.HelloResponse{Username: "lisi", Msg: "hello my man"}, nil
}

var (
	network = "tcp"
	addr    = "127.0.0.1:3001"
)

func main() {
	lis, err := net.Listen(network, addr)
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	pb.RegisterHelloServer(s, &Service{})
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}

}
