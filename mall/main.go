package main

import (
	"github.com/HsiaoCz/geek/mall/router"
	"github.com/gofiber/fiber/v2"
)

const (
	addr = "127.0.0.1:3303"
)

func main() {
	r := fiber.New()
	router.ResRouter(r)
	r.Listen(addr)
}
