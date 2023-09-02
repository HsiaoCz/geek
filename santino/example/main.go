package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	Addr = "127.0.0.1:9092"
)

func main() {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	RegisterRouter(r)
	srv := http.Server{
		Handler:      r,
		Addr:         Addr,
		WriteTimeout: 1500 * time.Millisecond,
		ReadTimeout:  1500 * time.Millisecond,
	}
	srv.ListenAndServe()
}
