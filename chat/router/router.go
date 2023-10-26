package router

import (
	"github.com/HsiaoCz/geek/chat/controller"
	"github.com/gofiber/fiber/v2"
)

func Router(r *fiber.App) {
	app := r.Group("/app")
	{
		// v1
		v1 := app.Group("/v1")
		{
			// auther api
			auth := v1.Group("/auth")
			{
				auth.Post("/register", controller.AuthRegister)
			}
		}
	}
}
