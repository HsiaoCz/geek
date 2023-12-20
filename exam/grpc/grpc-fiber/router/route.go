package router

import (
	"github.com/HsiaoCz/geek/exam/grpc/grpc-fiber/api"
	"github.com/gofiber/fiber/v2"
)

func Route() *fiber.App {
	r := fiber.New()
	app := r.Group("/app")
	{
		v1 := app.Group("/v1")
		{
			user := v1.Group("/user")
			{
				user.Post("/sinup", api.UserSinup)
				user.Post("/login", api.UserLogin)
			}
		}
	}
	return r
}
