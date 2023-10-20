package santino

import "net/http"

type Context struct {
	W http.ResponseWriter
	R *http.Request

	StatusCode int
	Path       string
	Method     string
}
