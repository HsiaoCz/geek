package service

import (
	"context"

	"github.com/HsiaoCz/geek/what/internal/data"
	"github.com/HsiaoCz/geek/what/internal/pb"
)

type UserService struct {
	pb.UnimplementedWhatServer
	ur data.UserRepo
}

func (u *UserService) UserSinup(ctx context.Context, in *pb.SinupRequest) (*pb.SinupResponse, error) {
	return &pb.SinupResponse{Msg: "注册成功", Code: 10000}, nil
}

func (u *UserService) UserLogin(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	return &pb.LoginResponse{Msg: "登录成功", Code: 10000}, nil
}

func New(ur data.UserRepo) *UserService {
	return &UserService{
		ur: ur,
	}
}
