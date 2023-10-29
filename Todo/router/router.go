package router

import (
	"github.com/HsiaoCz/geek/Todo/controller"
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
			}
		}
	}
}
