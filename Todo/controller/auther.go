package controller

import (
	"net/http"
	"strings"

	"github.com/HsiaoCz/geek/Todo/dao"
	"github.com/HsiaoCz/geek/Todo/utils"
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
	return dao.AuthReg(userR.Username, userR.Password, utils.GenIdentity(), userR.Email)

}

func AuthLogin(c *fiber.Ctx) error {
	userL := new(UserL)
	err := c.BodyParser(userL)
	if err != nil {
		return err
	}
	err = dao.AuthGetPasswd(userL.Username, userL.Password)
	if err != nil {
		return c.Status(http.StatusOK).JSON(fiber.Map{
			"Message": "用户名或密码不正确",
		})
	}
	return nil
}
