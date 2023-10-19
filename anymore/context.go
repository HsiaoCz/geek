package anymore

import (
	"encoding/json"
	"net/http"
)

type Context struct {
	w http.ResponseWriter
	r *http.Request
}

func (c *Context) JSON(code int, data any) error {
	c.r.Header.Set("Content-Type", "application/json")
	c.w.WriteHeader(code)
	return json.NewEncoder(c.w).Encode(data)
}
