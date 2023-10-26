package main

import (
	"github.com/HsiaoCz/geek/chat/router"
	"github.com/gofiber/fiber/v2"
)

func getfiber() *fiber.App {
	return fiber.New()
}

func main() {
	r := getfiber()
	router.Router(r)
	r.Listen("127.0.0.1:9900")
}
