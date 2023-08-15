package router

import (
	"github.com/HsiaoCz/geek/iml/controller"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine) {
	r.POST("api/v1/user/register", controller.UserRegister)
}
