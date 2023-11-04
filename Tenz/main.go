package main

import (
	"log"

	"github.com/HsiaoCz/geek/Tenz/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	r := fiber.New()
	router.RegisRoute(r)
	log.Fatal(r.Listen("127.0.0.1:9911"))
}
