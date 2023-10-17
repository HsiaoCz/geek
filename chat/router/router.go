package router

import "github.com/gofiber/fiber/v2"

func RsRouter(r *fiber.App){
	user:=r.Group("/api/v1/user")
	{
		user.Post("/register")
		user.Post("/login")
	}
}