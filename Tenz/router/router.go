package router

import (
	"github.com/HsiaoCz/geek/Tenz/controller"
	"github.com/HsiaoCz/geek/Tenz/middleware"
	"github.com/gofiber/fiber/v2"
)

func RegisRoute(r *fiber.App) {
	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			auther := v1.Group("/auther")
			{
				// 用户注册
				auther.Post("/register", controller.AutherRegister)
				// 用户登录
				auther.Post("/login", controller.AutherLogin)
				// 列出好友
				auther.Get("/listfriends", middleware.AutherVeryfi(), controller.AuthListFrident)
				// 修改个人信息
				auther.Put("/modify", middleware.AutherVeryfi(), controller.AuthModInfo)
			}
		}
	}
}
