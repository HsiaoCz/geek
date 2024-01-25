package route

import (
	"os"
	"time"

	"github.com/HsiaoCz/geek/what/api/av1"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Router(addr string) error {
	r := fiber.New()
	r.Use(logger.New(logger.Config{
		Next:         nil,
		Format:       "[${time}] ${status} - ${latency} ${method} ${path}\n",
		TimeFormat:   "2006/01/02 15:04:05",
		TimeZone:     "Local",
		TimeInterval: 500 * time.Millisecond,
		Output:       os.Stdout,
	}))
	app := r.Group("/app")
	{
		v1 := app.Group("/v1")
		{
			user := v1.Group("/user")
			{

				user.Post("/sinup", av1.UserSinup)
				user.Post("/login", av1.UserLogin)
				user.Get("/id", av1.GetUserByID)

			}
		}
	}
	return r.Listen(addr)
}
