package main

import (
	"net"
	"net/http"
	"os"
	"time"

	logsf "log"

	"github.com/HsiaoCz/geek/kit-grpc/pb"
	"github.com/HsiaoCz/geek/kit-grpc/services"
	"github.com/HsiaoCz/geek/kit-grpc/transports"
	"github.com/go-kit/log"
	"google.golang.org/grpc"
)

const (
	network  = "tcp"
	addrGRPC = "127.0.0.1:9901"
	addrHTTP = "127.0.0.1:9911"
)

func main() {
	// 这里启动http server
	svc := services.NewService()
	logger := log.NewLogfmtLogger(os.Stderr)
	r := transports.NewHTTPServer(svc, logger)
	srv := http.Server{
		Handler:      r,
		Addr:         addrHTTP,
		WriteTimeout: 1500 * time.Millisecond,
		ReadTimeout:  1500 * time.Millisecond,
	}
	go func() {
		logsf.Fatal(srv.ListenAndServe())
	}()
	lis, err := net.Listen(network, addrGRPC)
	if err != nil {
		logsf.Fatal(err)
	}
	s := grpc.NewServer()
	pb.RegisterHelloServiceServer(s, transports.NewGRPCServer(svc))
	go func() {
		logsf.Fatal(s.Serve(lis))
	}()

}
