package moo

import "net/http"

type Engine struct {
	srv    *http.Server
	stop   func() error
	router map[string]Handlefunc
}
