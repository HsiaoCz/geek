package middleware

import (
	"net/http"
	"strings"

	"github.com/HsiaoCz/geek/Todo/utils"
	"github.com/gofiber/fiber/v2"
)

func VerifyAuther() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		//token 可能在请求头,请求体，或url中
		// 这里假设放在请求头中
		autherHeader := c.Get("Authorization")
		if autherHeader == "" {
			return c.Status(http.StatusOK).JSON(fiber.Map{
				"Message": "请登录",
			})
		}
		parts := strings.SplitN(autherHeader, "", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			return c.Status(http.StatusOK).JSON(fiber.Map{
				"Message": "非法的token",
			})
		}
		mc, err := utils.ParseToken(parts[1])
		if err != nil {
			return c.Status(http.StatusOK).JSON(fiber.Map{
				"Message": "非法的token",
			})
		}
		// 将当前请求的UserId信息保存到请求的上下文中
		c.Locals("userID", mc.UserID)
		return c.Next()
	}
}
