package router

import (
	"github.com/HsiaoCz/geek/iml/controller"
	"github.com/gofiber/fiber/v2"
)

func RegisterRouter(r *fiber.App) {
	r.Post("api/v1/user/register", controller.UserRegister)
}
