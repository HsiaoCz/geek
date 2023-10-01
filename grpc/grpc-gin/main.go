package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/HsiaoCz/geek/grpc/grpc-gin/pb"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	net_work  = "tcp"
	gin_addr  = "127.0.0.1:3302"
	grpc_addr = "127.0.0.1:3303"
)

type Hello struct {
	Name string `json:"name"`
	Msg  string `json:"msg"`
}

// grpc 结合gin
// 或者别的web框架也行

type Service struct {
	pb.UnimplementedGreeterServer
}

func (s *Service) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	fmt.Println(in.GetName(), ":", in.GetMsg())
	return &pb.HelloResponse{Name: "lisi", Msg: "Hello what's up"}, nil
}

func main() {
	// 这里应该写在grpc的服务端，这里不应该写在一起
	lis, err := net.Listen(net_work, grpc_addr)
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &Service{})
	err = s.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
	r := gin.Default()
	r.POST("/api/hello", handleHello)
	r.Run(gin_addr)
}

func handleHello(c *gin.Context) {
	hello := new(Hello)
	c.BindJSON(hello)
	conn, err := grpc.Dial(grpc_addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	cc := pb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := cc.SayHello(ctx, &pb.HelloRequest{Name: hello.Name, Msg: hello.Msg})
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"name": res.GetName(),
		"msg":  res.GetMsg(),
	})
}
