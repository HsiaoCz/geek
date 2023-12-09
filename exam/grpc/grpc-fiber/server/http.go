package server

import (
	"github.com/HsiaoCz/geek/exam/grpc/grpc-fiber/router"
)

func RegisterHttpServer(addr string) error {
	r := router.Route()
	return r.Listen(addr)
}
