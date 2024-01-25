package slick

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type ErrorHandler func(error, *Context) error
type Handler func(c *Context) error

type Slick struct {
	ErrorHandler ErrorHandler
	router       *httprouter.Router
}

type Context struct {
	response http.ResponseWriter
	request  *http.Request
	ctx      context.Context
}

func New() *Slick {
	return &Slick{
		router:       httprouter.New(),
		ErrorHandler: defaultErrorHandler,
	}
}

func (s *Slick) Get(path string, h Handler, plugs ...Handler) error {
	s.router.GET(path, s.makeHTTPRouterHandler(h))
	return nil
}

func (s *Slick) makeHTTPRouterHandler(h Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		ctx := &Context{
			response: w,
			request:  r,
			ctx:      context.Background(),
		}
		if err := h(ctx); err != nil {
			s.ErrorHandler(err, ctx)
		}
	}
}

func (s *Slick) Start(addr string) error {
	return http.ListenAndServe(addr, s.router)
}

func defaultErrorHandler(err error, c *Context) error {
	slog.Error("error", "err", err)
	return nil
}
