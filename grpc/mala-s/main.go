package main

import (
	"context"
	"log"
	"math/rand"
	"net"
	"time"

	"github.com/HsiaoCz/geek/grpc/mala-s/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type Service struct {
	pb.UnimplementedGreeterServer
}

func (s *Service) UserLogin(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.DataLoss, "UserLogin: need the metadata")
	}
	if t, ok := md["token"]; ok {
		if len(t) < 1 || t[0] != "app-test-hsiaocz" {
			return nil, status.Errorf(codes.Unauthenticated, "认证失败")
		}
	}
	if in.GetPasswrod() != in.GetRepassword() {
		return nil, status.Errorf(codes.OK, "请检查用户名和密码")
	}
	return &pb.LoginResponse{Identity: GenIdentity(), Msg: "登录成功"}, nil
}

func GenIdentity() int64 {
	randm := rand.New(rand.NewSource(time.Now().UnixNano()))
	return randm.Int63n(1000000)
}

const (
	network = "tcp"
	addr    = "127.0.0.1:9911"
)

func main() {
	lis, err := net.Listen(network, addr)
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &Service{})
	err = s.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}
