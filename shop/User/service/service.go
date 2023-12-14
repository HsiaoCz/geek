package service

import "github.com/HsiaoCz/geek/shop/User/pb/uv1"

type Service struct {
	uv1.UnimplementedUserServiceServer
}

func NewService() *Service {
	return &Service{}
}
