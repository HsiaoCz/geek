package server

import (
	"net"

	"github.com/HsiaoCz/geek/shop/User/pb/uv1"
	"github.com/HsiaoCz/geek/shop/User/service"
	"google.golang.org/grpc"
)

func ResGrpc(network string, addr string) error {
	listen, err := net.Listen(network, addr)
	if err != nil {
		return err
	}
	server := grpc.NewServer()
	uv1.RegisterUserServiceServer(server, service.NewService())
	return server.Serve(listen)
}
