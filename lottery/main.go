package main

import (
	"log"
	"net/http"
	"time"

	"github.com/HsiaoCz/geek/lottery/config"
	"github.com/HsiaoCz/geek/lottery/dao"
	"github.com/HsiaoCz/geek/lottery/router"
	"github.com/gin-gonic/gin"
)

const (
	addr = "127.0.0.1:9901"
)

func main() {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	err := config.Init()
	if err != nil {
		log.Println(err)
		return
	}
	err = dao.InitMysql()
	if err != nil {
		log.Println(err)
		return
	}
	router.RegisterRouter(r)
	srv := http.Server{
		Handler:      r,
		Addr:         addr,
		ReadTimeout:  1500 * time.Millisecond,
		WriteTimeout: 1500 * time.Millisecond,
	}
	log.Fatal(srv.ListenAndServe())
}
