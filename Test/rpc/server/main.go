package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

// rpc server

type Args struct {
	X int
	Y int
}

// ServiceA 自定义一个结构体类型

type ServiceA struct{}

func (s *ServiceA) Add(args *Args, reply *int) error {
	*reply = args.X + args.Y
	return nil
}

func main() {
	service := new(ServiceA)
	rpc.Register(service) // 注册rpc
	rpc.HandleHTTP()      // 基于HTTP协议
	l, e := net.Listen("tcp", ":9091")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	http.Serve(l, nil)

	// TCP版的调用
	// service := new(ServiceA)
	// rpc.Register(service) // 注册RPC服务
	// l, e := net.Listen("tcp", ":9091")
	// if e != nil {
	// 	log.Fatal("listen error:", e)
	// }
	// for {
	// 	conn, _ := l.Accept()
	// 	rpc.ServeConn(conn)
	// }
}
