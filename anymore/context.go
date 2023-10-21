package anymore

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Context struct {
	W          http.ResponseWriter
	R          *http.Request
	Path       string
	Method     string
	StatusCode int
	Params     map[string]string
}

func newContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		W:      w,
		R:      r,
		Path:   r.URL.Path,
		Method: r.Method,
	}
}

func (c *Context) PostForm(key string) string {
	return c.R.FormValue(key)
}

func (c *Context) Query(key string) string {
	return c.R.URL.Query().Get(key)
}

func (c *Context) Status(code int) {
	c.StatusCode = code
	c.W.WriteHeader(code)
}

func (c *Context) SetHeader(key string, value string) {
	c.W.Header().Set(key, value)
}

func (c *Context) String(code int, format string, data ...any) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.W.Write([]byte(fmt.Sprintf(format, data...)))
}

func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.W.Write(data)
}

func (c *Context) HTML(code int, data []byte) {
	c.Status(code)
	c.W.Write(data)
}

func (c *Context) JSON(code int, data any) error {
	c.R.Header.Set("Content-Type", "application/json")
	c.W.WriteHeader(code)
	return json.NewEncoder(c.W).Encode(data)
}

func (c *Context) Param(key string) string {
	return c.Params[key]
}
