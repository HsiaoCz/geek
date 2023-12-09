package server

import (
	"net"

	"google.golang.org/grpc"
)

func RegGrpc(network string, addr string) error {
	listen, err := net.Listen(network, addr)
	if err != nil {
		return err
	}
	s := grpc.NewServer()
	return s.Serve(listen)
}
