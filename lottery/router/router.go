package router

import (
	"github.com/HsiaoCz/geek/lottery/controller"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine) {
	r.POST("/user/register", controller.UserRegister)
	r.POST("/user/login", controller.UserLogin)
	r.GET("/user/lottery", controller.UserGetLottery)
}
