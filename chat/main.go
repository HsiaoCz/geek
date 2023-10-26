package main

import (
	"github.com/HsiaoCz/geek/chat/router"
	"github.com/gofiber/fiber/v2"
)

func getfiber() *fiber.App {
	return fiber.New()
}

func main() {
	router.Router(getfiber())
}
