package main

import (
	"net/http"

	"github.com/HsiaoCz/geek/kit/endpoints"
	"github.com/HsiaoCz/geek/kit/service"
	"github.com/HsiaoCz/geek/kit/transport"
	httrans "github.com/go-kit/kit/transport/http"
)

const (
	addr = "127.0.0.1:9911"
)

func main() {
	user := &service.UserService{}
	endp := endpoints.GetUserEndpoint(user)

	// 使用go-kit的api创建handler
	srvhandler := httrans.NewServer(endp, transport.DecodeUserRequest, transport.EncodeUserResponse)
	http.ListenAndServe(addr, srvhandler)
}

