package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	r := fiber.New()
	log.Fatal(r.Listen("127.0.0.1:9911"))
}
