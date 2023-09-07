package main

import (
	"log"
	"net/http"
	"time"

	"github.com/HsiaoCz/geek/iml/router"
	"github.com/gin-gonic/gin"
)

const (
	addr = "127.0.0.1:9091"
)

func main() {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	router.RegisterRouter(r)
	srv := http.Server{
		Handler:      r,
		Addr:         addr,
		WriteTimeout: 1500 * time.Millisecond,
		ReadTimeout:  1500 * time.Millisecond,
	}
	log.Fatal(srv.ListenAndServe())
}
