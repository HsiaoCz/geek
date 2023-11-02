package router

import (
	"github.com/HsiaoCz/geek/Todo/controller"
	"github.com/HsiaoCz/geek/Todo/middleware"
	"github.com/gofiber/fiber/v2"
)

func ResRoute(r *fiber.App) {
	app := r.Group("/app")
	{
		v1 := app.Group("/v1")
		{
			auth := v1.Group("/auther")
			{
				auth.Post("/register", controller.AuthRegister)
				auth.Post("/login", controller.AuthLogin)
				auth.Get("/todolist", controller.AuthGetTodoList, middleware.VerifyAuther())
			}
		}
	}
}
