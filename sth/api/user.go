package api

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type user struct {
}

func newUser() *user {
	return &user{}
}

func (u *user) handleUserRegister(c *fiber.Ctx) error {
	userR := new(UserR)
	err := c.BodyParser(userR)
	if err != nil {
		return err
	}
	if userR.Password != userR.RePassword {
		return c.Status(http.StatusOK).JSON(fiber.Map{
			"Message": "请检查用户名或密码",
		})
	}
	return nil
}
