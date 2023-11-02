package main

import (
	"log"
	"net/http"
	"time"

	"github.com/HsiaoCz/geek/Todo/config"
	"github.com/HsiaoCz/geek/Todo/dao"
	"github.com/HsiaoCz/geek/Todo/router"
	"github.com/gin-gonic/gin"
)

func main() {
	err := config.InitConf()
	if err != nil {
		log.Fatal(err)
	}
	err = dao.InitMysql()
	if err != nil {
		log.Fatal(err)
	}
	r := gin.Default()
	r.Use(gin.Logger(), gin.Recovery())
	router.ResRoute(r)
	srv := http.Server{
		Handler:      r,
		Addr:         config.Conf.AC.Port,
		ReadTimeout:  1500 * time.Millisecond,
		WriteTimeout: 1500 * time.Millisecond,
	}
	log.Fatal(srv.ListenAndServe())
}
