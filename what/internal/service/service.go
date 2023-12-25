package service

import (
	"context"

	"github.com/HsiaoCz/geek/what/internal/pb"
)

type UserCase struct {
	pb.UnimplementedWhatServer
}

func (u *UserCase) UserSinup(ctx context.Context, in *pb.SinupRequest) (*pb.SinupResponse, error) {
	return &pb.SinupResponse{Msg: "注册成功", Code: 10000}, nil
}

func (u *UserCase) UserLogin(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	return &pb.LoginResponse{Msg: "登录成功", Code: 10000}, nil
}

func New() *UserCase {
	return &UserCase{}
}
