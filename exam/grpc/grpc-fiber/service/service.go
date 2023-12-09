package service

import (
	"context"

	"github.com/HsiaoCz/geek/exam/grpc/grpc-fiber/pb"
)

type Service struct {
	pb.UnimplementedHelloServer
}

func (s *Service) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	return &pb.LoginResponse{Msg: "登录成功", Code: 10000}, nil
}

func (s *Service) Sinup(ctx context.Context, in *pb.SinupRequest) (*pb.SinupResponse, error) {
	return &pb.SinupResponse{Msg: "注册成功", Code: 10000}, nil
}

func NewService() *Service {
	return &Service{}
}
