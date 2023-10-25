package moo

import "net/http"

// 引擎需要具备的方法
// 这里http.Handler是必须的
// 启动方法，和关闭方法
// 还需要路由注册方法
type server interface {
	http.Handler
	Start(addr string) error
	Stop() error
	addRouter(method string, pattern string, handlefunc Handlefunc)
}
