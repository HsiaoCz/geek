package router

import (
	"github.com/HsiaoCz/geek/disk/controller"
	"github.com/gofiber/fiber/v2"
)

func Route(r *fiber.App) {
	app := r.Group("/app")
	{
		v1 := app.Group("/v1")
		{
			file := v1.Group("/file")
			{
				file.Post("/upload", controller.UploadFile)
			}
		}
	}
}
