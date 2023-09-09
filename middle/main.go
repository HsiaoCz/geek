package main

import (
	"log"

	"github.com/HsiaoCz/geek/middle/dao/mysql"
	"github.com/HsiaoCz/geek/middle/router"
)

const (
	addr = "127.0.0.1:9091"
)

func main() {
	router.RegisterRouter(addr)
	if err := mysql.InitMysql(); err != nil {
		log.Println(err)
		return
	}
}
