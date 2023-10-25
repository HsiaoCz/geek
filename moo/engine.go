package moo

import (
	"log"
	"net/http"
)

type Engine struct {
	srv    *http.Server
	stop   func() error
	router *router
}

// ServeHTTP 实现http.Handler接口
// 使得Engine成为一个handler
func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 匹配路由
	n, params, ok := e.router.getRouter(r.Method, r.URL.Path)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 NOT FOUND"))
		return
	}

	// 构造当前请求的上下文
	c := NewContext(w, r)
	c.params = params
	log.Printf("request %s-%s", c.Method, c.Pattern)
	n.handlefunc(c)
}

// 启动服务
// 这里还可以使用选项模式给启动添加一些选项，比如ReadtimeOut和WriteTimeOut

func (e *Engine) Start(addr string) error {
	e.srv = &http.Server{
		Addr:    addr,
		Handler: e,
	}
	log.Printf("the server is running on port%s\n", addr)
	return e.srv.ListenAndServe()
}

// 关闭服务
func (e *Engine) Stop() error {
	return e.stop()
}

// 请求方法篇
// 请求方法后面还需要添加中间件
func (e *Engine) GET(pattern string, hanlerfunc Handlefunc) {
	e.router.addRouter(http.MethodGet, pattern, hanlerfunc)
}
func (e *Engine) POST(pattern string, hanlerfunc Handlefunc) {
	e.router.addRouter(http.MethodPost, pattern, hanlerfunc)
}
func (e *Engine) PATCH(pattern string, hanlerfunc Handlefunc) {
	e.router.addRouter(http.MethodPatch, pattern, hanlerfunc)
}
func (e *Engine) PUT(pattern string, hanlerfunc Handlefunc) {
	e.router.addRouter(http.MethodPut, pattern, hanlerfunc)
}
func (e *Engine) DELETE(pattern string, hanlerfunc Handlefunc) {
	e.router.addRouter(http.MethodDelete, pattern, hanlerfunc)
}

func (e *Engine) TRACE(pattern string, hanlerfunc Handlefunc) {
	e.router.addRouter(http.MethodTrace, pattern, hanlerfunc)
}
func (e *Engine) OPTIONS(pattern string, hanlerfunc Handlefunc) {
	e.router.addRouter(http.MethodOptions, pattern, hanlerfunc)
}
func (e *Engine) CONNECT(pattern string, hanlerfunc Handlefunc) {
	e.router.addRouter(http.MethodConnect, pattern, hanlerfunc)
}
func (e *Engine) HEAD(pattern string, handlefunc Handlefunc) {
	e.router.addRouter(http.MethodHead, pattern, handlefunc)
}
