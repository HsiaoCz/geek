package anymore

import (
	"fmt"
	"net/http"
)

type Engine struct {
	router map[string]HandleFunc
}

func App() *Engine {
	return &Engine{
		router: make(map[string]HandleFunc),
	}
}

// addRouter 添加路由的方法
// 不过这个后面会拆分的
func (e *Engine) addRouter(method string, pattern string, handler HandleFunc) {
	key := method + "-" + pattern
	e.router[key] = handler
}

// 引擎上的路由注册方法
func (e *Engine) GET(pattern string, handler HandleFunc) {
	e.addRouter("GET", pattern, handler)
}
func (e *Engine) POST(pattern string, handler HandleFunc) {
	e.addRouter("POST", pattern, handler)
}
func (e *Engine) PUT(pattern string, handler HandleFunc) {
	e.addRouter("PUT", pattern, handler)
}
func (e *Engine) DELETE(pattern string, handler HandleFunc) {
	e.addRouter("DELETE", pattern, handler)
}
func (e *Engine) HEAD(pattern string, handler HandleFunc) {
	e.addRouter("HEAD", pattern, handler)
}
func (e *Engine) TRANCE(pattern string, handler HandleFunc) {
	e.addRouter("TRANCE", pattern, handler)
}

func (e *Engine) OPTIONS(pattern string, handler HandleFunc) {
	e.addRouter("OPTIONS", pattern, handler)
}
func (e *Engine) PATCH(pattern string, handler HandleFunc) {
	e.addRouter("PATCH", pattern, handler)
}
func (e *Engine) CONNECT(pattern string, handler HandleFunc) {
	e.addRouter("CONNECT", pattern, handler)
}

// ServeHTTP 这个方法是引擎必须实现的
func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := r.Method + "-" + r.URL.Path
	if handler, ok := e.router[key]; ok {
		handler(w, r)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND:%s\n", r.URL)
	}
}

// Listen
func (e *Engine) Listen(addr string) error {
	return http.ListenAndServe(addr, e)
}
