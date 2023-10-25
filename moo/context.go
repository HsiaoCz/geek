package moo

import "net/http"

type Context struct {
	// Response
	W http.ResponseWriter
	// Request
	R *http.Request
	// Method
	Method string
	// request URL
	Pattern string
	// router params
	params map[string]string
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		W:       w,
		R:       r,
		Method:  r.Method,
		Pattern: r.URL.Path,
	}
}
