package router

import (
	"github.com/HsiaoCz/geek/chat/controller"
	"github.com/gofiber/fiber/v2"
)

func Router(r *fiber.App) {
	v1 := r.Group("/api")
	{
		auth := v1.Group("/v1")
		{
			auth.Post("/register", controller.AuthRegister)
		}
	}
}
