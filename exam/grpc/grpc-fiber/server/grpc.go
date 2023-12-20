package server

import (
	"net"

	"github.com/HsiaoCz/geek/exam/grpc/grpc-fiber/pb"
	"github.com/HsiaoCz/geek/exam/grpc/grpc-fiber/service"
	"google.golang.org/grpc"
)

func RegisterGrpcServer(net_work string, addr string) error {
	listen, err := net.Listen(net_work, addr)
	if err != nil {
		return err
	}
	s := grpc.NewServer()
	pb.RegisterHelloServer(s, service.NewService())
	return s.Serve(listen)
}
