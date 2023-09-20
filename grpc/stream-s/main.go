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

	"github.com/HsiaoCz/geek/grpc/stream-s/pb"
	"google.golang.org/grpc"
)

type Service struct {
	pb.UnimplementedGreeterServer
}

func (s *Service) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	if in.Password != in.Repasswrod {
		return &pb.LoginResponse{Msg: "请确认用户名或密码"}, nil
	}
	return &pb.LoginResponse{Msg: "登录成功"}, nil
}
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
func (s *Service) HandleHi(stream pb.Greeter_HandleHiServer) error {
	for {
		res, err := stream.Recv()
		if err != nil {
			log.Fatal(err)
		}
		if err == io.EOF {
			return stream.SendAndClose(&pb.HiResponse{
				Something: "你好,my man",
			})
		}
		fmt.Println(res.GetContent())
	}
}
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
		fmt.Println(res.GetChatcontent())
	}
}

func GenPID() string {
	randm := rand.New(rand.NewSource(time.Now().UnixNano()))
	return strconv.Itoa(randm.Intn(100000))
}

const (
	network = "tcp"
	addr    = "127.0.0.1:9922"
)

func main() {
	lis, err := net.Listen(network, addr)
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()                   // 创建grpc服务
	pb.RegisterGreeterServer(s, &Service{}) // 注册服务
	err = s.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}
