package middleware

import (
	"net/http"
	"strings"

	"github.com/HsiaoCz/geek/Tenz/utils"
	"github.com/gofiber/fiber/v2"
)

func AutherVeryfi() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		// 客户端携带Token有三种方式:1.放在请求头 2.放在请求体
		// 3 .放在URL
		// 这里假设Token放在Header的Autherization中，并使用Bearer开头
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"Message": "请登录",
			})
		}
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) == 2 && parts[0] == "Bearer" {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"Message": "无效的Token",
			})
		}
		// parts[1]是获取到的tokenString
		mc, err := utils.ParseToken(parts[1])
		if err != nil {
			return c.Status(http.StatusBadGateway).JSON(fiber.Map{
				"Message": "无效的Token",
			})
		}
		// 将当前请求的UserID信息保存到请求的上下文C上
		c.Locals("userID", mc.UserID)
		// 后续的处理函数
		return c.Next()
	}
}
