package santino

import "net/http"

// 我们开始实现一个自己的框架
// 第一步，实现一个自己的Handler
// 这个santino 类似于gin里面engine

type ISantino interface {
	// 首先我们的引擎必须实现这个接口才能注册转发路由
	http.Handler

	// 服务启动的方法
	// 服务启动需要知道address
	Start(addr string) error

	// 服务关闭的方法
	Stop() error

	// 添加路由的方法
	// 将方法，path,和具体的方法注册进路由
	addRouter(method string, pattern string, handlefunc HandleFunc)
}

type Santino struct {
	srv  *http.Server
	stop func() error
	*router
}

// http.ListenAndServe(addr,handler)
// 第二个参数handler本质是一个接口
// 这个接口又一个方法
// 就是ServeHTTP(http.ResponseWriter,*http.Request)
// 我们实现这个接口，就有了自己的handler
// 这里，这个ServeHTTP作用是什么呢？
// 简单说，将一个请求匹配到注册到这个路由上的HandleFunc上
func (s *Santino) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}

// 这里的new方法类似于gin.New()返回一个默认的Santino
func New() *Santino {
	return &Santino{}
}

func (s *Santino) GET(pattern string, handlefunc HandleFunc)     {}
func (s *Santino) POST(pattern string, handlefunc HandleFunc)    {}
func (s *Santino) PUT(pattern string, handlefunc HandleFunc)     {}
func (s *Santino) DELETE(pattern string, handlefunc HandleFunc)  {}
func (s *Santino) HEAD(pattern string, handlefunc HandleFunc)    {}
func (s *Santino) TRACE(pattern string, handlefunc HandleFunc)   {}
func (s *Santino) OPTIONS(pattern string, handlefunc HandleFunc) {}
func (s *Santino) PATCH(pattern string, handlefunc HandleFunc)   {}
func (s *Santino) CONNECT(pattern string, handlefunc HandleFunc) {}

// 匹配没匹配到的
func (s *Santino) Any(pattern string, handlefunc HandleFunc) {}
