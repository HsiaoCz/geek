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
			todo := v1.Group("/todo", middleware.VerifyAuther())
			{
				todo.POST("/add", controller.AddTodoList)
				todo.PUT("/mod", controller.ModTodoList)
				todo.GET("/get", controller.GetTodoList)
				todo.DELETE("/delete", controller.DeleteTodoList)
				todo.GET("/complete", controller.CompleteTodoList)
				todo.DELETE("/clean", controller.CleanTodoList)
			}
		}
	}
}
