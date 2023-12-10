package service

import (
	"context"

	"github.com/HsiaoCz/geek/what/internal/pb/gv1"
)

type UserCase struct {
	gv1.UnimplementedWhatUserPartServer
}

func (u *UserCase) UserSinup(ctx context.Context, in *gv1.SinupRequest) (*gv1.SinupResponse, error) {
	return &gv1.SinupResponse{Msg: "注册成功", Code: 10000}, nil
}

func (u *UserCase) UserLogin(ctx context.Context, in *gv1.LoginRequest) (*gv1.LoginResponse, error) {
	return &gv1.LoginResponse{Msg: "登录成功", Code: 10000}, nil
}

func New() *UserCase {
	return &UserCase{}
}
