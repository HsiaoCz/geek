package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"strconv"
	"time"

	"github.com/HsiaoCz/geek/exam/grpc/stream-s/pb"
	"google.golang.org/grpc"
)

type Service struct {
	pb.UnimplementedGreeterServer
}

func (s *Service) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	if in.GetPassword() != in.GetRepasswrod() {
		return &pb.LoginResponse{Msg: "请确认密码"}, nil
	}
	return &pb.LoginResponse{Msg: "登录成功"}, nil
}

// 服务端流式rpc
func (s *Service) Hello(in *pb.HelloRequest, stream pb.Greeter_HelloServer) error {
	fmt.Println(in.Content)
	words := []string{"你好", "hello", "halo", "sawadika", "bonjor"}
	for _, word := range words {
		data := &pb.HelloResponse{Something: word}
		if err := stream.Send(data); err != nil {
			return err
		}
	}
	return nil
}

// 客户端流式grpc

func (s *Service) HandleHi(stream pb.Greeter_HandleHiServer) error {
	for {
		res, err := stream.Recv()
		if err != nil {
			log.Fatal(err)
		}
		if err == io.EOF {
			return stream.SendAndClose(&pb.HiResponse{
				Something: "你好,man",
			})
		}
		fmt.Println(res.GetContent())
	}
}

// 双向流式grpc
func (s *Service) HandleChat(stream pb.Greeter_HandleChatServer) error {
	for {
		res, err := stream.Recv()
		if err != nil {
			log.Fatal(err)
		}
		if err == io.EOF {
			return stream.Send(&pb.ChatResponse{
				Anwser: "你好:" + res.GetName(),
				Pid:    GenPID(),
			})
		}
		fmt.Println()
	}
}

func GenPID() string {
	randm := rand.New(rand.NewSource(time.Now().UnixNano()))
	return strconv.Itoa(randm.Intn(100000))
}

var (
	network = "tcp"
	addr    = "127.0.0.1:9001"
)

func main() {
	listen, err := net.Listen(network, addr)
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &Service{})
	if err := s.Serve(listen); err != nil {
		log.Fatal(err)
	}
}
