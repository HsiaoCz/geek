package route

import "github.com/gofiber/fiber/v2"

func Router(addr string) error {
	r := fiber.New()
	app := r.Group("/app")
	{
		v1 := app.Group("/v1")
		{
			user := v1.Group("/user")
			{
				user.Post("/signup")
				user.Post("/login")
			}
		}
	}
	return r.Listen(addr)
}
