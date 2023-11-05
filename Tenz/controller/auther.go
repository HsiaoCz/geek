package controller

import (
	"net/http"

	"github.com/HsiaoCz/geek/Tenz/dao"
	"github.com/gofiber/fiber/v2"
)

func AutherRegister(c *fiber.Ctx) error {
	userR := new(UserR)
	// 这里序列化传递来的数据
	if err := c.BodyParser(userR); err != nil {
		return err
	}
	// 这里判断传递过来的username和passwrod和RePassword是否为空
	if len(userR.Username) == 0 || len(userR.Password) == 0 || len(userR.RePassword) == 0 {
		return c.Status(http.StatusOK).JSON(fiber.Map{
			"Message": "请输入用户名或密码",
		})
	}
	// 密码和确认密码不相等的话
	if userR.Password != userR.RePassword {
		c.Status(http.StatusOK).JSON(fiber.Map{
			"Message": "请确认密码",
		})
	}
	// 这里看一下手机号是否已经注册了
	if result := dao.AuthTestPhoneNumber(userR.PhoneNumber); result > 0 {
		return c.Status(http.StatusOK).JSON(fiber.Map{
			"Message": "当前手机已经注册",
		})
	}

	// 将用户信息入库
	if err := dao.AuthCreate(userR.Username, userR.Password, userR.PhoneNumber); err != nil {
		return err
	}
	// 返回注册成功
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"Message": "注册成功",
	})
}

func AutherLogin(c *fiber.Ctx) error {
	return nil
}
