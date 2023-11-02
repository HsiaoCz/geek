package router

import (
	"github.com/HsiaoCz/geek/Todo/controller"
	"github.com/HsiaoCz/geek/Todo/middleware"
	"github.com/gin-gonic/gin"
)

func ResRoute(r *gin.Engine) {
	app := r.Group("/app")
	{
		v1 := app.Group("/v1")
		{
			auth := v1.Group("/auther")
			{
				auth.POST("/register", controller.AuthRegister)
				auth.POST("/login", controller.AuthLogin)
				auth.GET("/todolist", middleware.VerifyAuther(), controller.AuthGetTodoList)
			}
		}
	}
}
