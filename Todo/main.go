package main

import (
	"log"

	"github.com/HsiaoCz/geek/Todo/config"
	"github.com/HsiaoCz/geek/Todo/dao"
	"github.com/HsiaoCz/geek/Todo/router"
	"github.com/gofiber/fiber/v2"
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
	r := fiber.New()
	router.ResRoute(r)
	log.Fatal(r.Listen("127.0.0.1:8090"))
}
