package server

import (
	"net"

	"github.com/HsiaoCz/geek/what/internal/pb"
	"github.com/HsiaoCz/geek/what/internal/service"
	"google.golang.org/grpc"
)

func RegGrpc(network string, addr string) error {
	listen, err := net.Listen(network, addr)
	if err != nil {
		return err
	}
	s := grpc.NewServer()
	pb.RegisterWhatServer(s, service.New())
	return s.Serve(listen)
}
