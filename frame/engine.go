package frame

import "net/http"

type Engine struct {
	srv  http.Server
	addr string
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {}
func (e *Engine) Start() error {
	e.srv.Addr = e.addr
	return e.srv.ListenAndServe()
}
func (e *Engine) GET(pattren string, handler Handlefunc) {}

func (e *Engine) POST(pattren string, handler Handlefunc) {}

func (e *Engine) PUT(pattren string, handler Handlefunc) {}

func (e *Engine) DELETE(pattren string, handler Handlefunc) {}

func (e *Engine) HEAD(pattren string, handler Handlefunc) {}
