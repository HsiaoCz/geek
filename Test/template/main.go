package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/app/v1/auther/login", HandleUserLogin)
	log.Fatal(r.Run("127.0.0.1:3002"))
}

func HandleUserLogin(c *gin.Context) {
	
}
