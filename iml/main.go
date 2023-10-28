package main

import (
	"log"

	"github.com/HsiaoCz/geek/iml/router"
	"github.com/gofiber/fiber/v2"
)

const (
	addr = "127.0.0.1:9091"
)

func main() {
	r := fiber.New()
	router.RegisterRouter(r)
	log.Fatal(r.Listen(addr))
}
