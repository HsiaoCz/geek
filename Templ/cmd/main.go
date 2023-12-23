package main

import (
	"log"

	"github.com/HsiaoCz/geek/Templ/handler"
	"github.com/labstack/echo/v4"
)

var (
	addr = "127.0.0.1:3001"
)

func main() {
	app := echo.New()
	app.GET("/api/:username", handler.Hello)
	log.Fatal(app.Start(addr))
}
