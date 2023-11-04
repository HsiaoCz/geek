package router

import (
	"github.com/HsiaoCz/geek/Tenz/controller"
	"github.com/gofiber/fiber/v2"
)

func RegisRoute(r *fiber.App) {
	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			auther := v1.Group("/auther")
			{
				auther.Post("/register", controller.AutherRegister)
				auther.Post("/login", controller.AutherLogin)
			}
		}
	}
}
