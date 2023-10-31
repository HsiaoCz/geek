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
	if !strings.Contains(userR.Email, "@") {
		return c.Status(http.StatusOK).JSON("请输入合法的邮箱")
	}
	if result := dao.AuthGetUserByUsernameAndEmail(userR.Username, userR.Email); result > 0 {
		return c.Status(http.StatusOK).JSON("该用户当前已经存在")
	}
	if err := dao.AuthReg(userR.Username, userR.Password, utils.GenIdentity(), userR.Email); err == nil {
		return c.Status(http.StatusOK).JSON("注册成功")
	}
	return err
}

func AuthLogin(c *fiber.Ctx) error {
	userL := new(UserL)
	err := c.BodyParser(userL)
	if err != nil {
		return err
	}
	if result := dao.AuthGetPasswdAndEmial(userL.Username, userL.Password, userL.Email); result == 0 {
		return c.Status(http.StatusOK).JSON(fiber.Map{
			"Message": "用户名或密码不正确",
		})
	}
	return c.Status(http.StatusOK).JSON("登录成功!")
}
