package main

import (
	"log"

	"github.com/HsiaoCz/geek/img/dao/mysql"
	"github.com/HsiaoCz/geek/img/router"
)

const (
	addr = "127.0.0.1:9991"
)

func main() {
	if err := mysql.ConnMysql(); err != nil {
		log.Fatal(err)
	}
	router.RegRouter(addr)
}
