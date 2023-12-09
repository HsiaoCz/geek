package route

import (
	"github.com/HsiaoCz/geek/what/api"
	"github.com/gofiber/fiber/v2"
)

func Router(addr string) error {
	r := fiber.New()
	app := r.Group("/app")
	{
		v1 := app.Group("/v1")
		{
			user := v1.Group("/user")
			{

				user.Post("/sinup", api.UserSinup)
				user.Post("/login", api.UserSinup)

			}
		}
	}
	return r.Listen(addr)
}
