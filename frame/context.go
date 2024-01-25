package frame

import "net/http"

type Context struct {
	r *http.Request
	w http.ResponseWriter
}

func (c *Context) String(code int, v string) {}
func (c *Context) JSON(code int, v any)      {}
func (c *Context) HTML()                     {}

