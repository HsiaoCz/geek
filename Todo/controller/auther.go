package controller

import (
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func AuthRegister(c *fiber.Ctx) error {
	userR := new(UserR)
	err := c.BodyParser(userR)
	if err != nil {
		return err
	}
	if len(userR.Username) == 0 || len(userR.Password) == 0 {
		return c.Status(http.StatusOK).JSON("用户名或密码不能为空")
	}
	if userR.Password != userR.RePassword {
		return c.Status(http.StatusOK).JSON("请确认密码")
	}
	if strings.Contains(userR.Email, "@") {
		return c.Status(http.StatusOK).JSON("请输入合法的邮箱")
	}
	return nil
}
