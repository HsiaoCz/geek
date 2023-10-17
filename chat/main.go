package main

import (
	"log"

	"github.com/HsiaoCz/geek/chat/conf"
	"github.com/gofiber/fiber/v2"
)

func main() {
	err := conf.Init()
	if err != nil {
		log.Fatal(err)
	}
	r := fiber.New()
	r.Listen(conf.Conf.App.AppAddr)
}
