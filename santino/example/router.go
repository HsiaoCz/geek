package main

import "github.com/gin-gonic/gin"

func RegisterRouter(r *gin.Engine) {
	r.POST("/user/register", HandleUserRegister)
}
