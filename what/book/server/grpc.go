package server

import (
	"net"

	"github.com/HsiaoCz/geek/what/book/pb"
	"github.com/HsiaoCz/geek/what/book/service"
	"google.golang.org/grpc"
)

func RegisterGrpc(addr string, network string) error {
	conn, err := net.Listen(network, addr)
	if err != nil {
		return err
	}
	server := grpc.NewServer()
	pb.RegisterBookServiceServer(server, service.NewBookService())
	return server.Serve(conn)
}
