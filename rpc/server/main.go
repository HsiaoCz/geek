package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Args struct {
	X int
	Y int
}

type Service struct{}

func (s *Service) Add(args *Args, reply *int) error {
	*reply = args.X + args.Y
	return nil
}

const (
	network = "tcp"
	addr    = "127.0.0.1:9991"
)

func main() {
	service := new(Service)
	rpc.Register(service) // 注册rpc服务
	rpc.HandleHTTP()      // 基于HTTP协议

	l, err := net.Listen(network, addr)
	if err != nil {
		log.Fatal(err)
	}
	http.Serve(l, nil)
}
