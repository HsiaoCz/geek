package main

import (
	"github.com/HsiaoCz/geek/iml/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	router.RegisterRouter(r)
}
