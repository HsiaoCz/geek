package controller

import (
	"net/http"
	"strconv"

	"github.com/HsiaoCz/geek/Tenz/dao"
	"github.com/HsiaoCz/geek/Tenz/utils"
	"github.com/gofiber/fiber/v2"
)

func AutherRegister(c *fiber.Ctx) error {
	userR := new(UserR)
	// 这里序列化传递来的数据
	if err := c.BodyParser(userR); err != nil {
		return err
	}
	// 这里判断传递过来的username和passwrod和RePassword是否为空
	// 判空操作其实可以放在json的tag里面
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
	// 拿到用户信息
	userL := new(UserL)
	if err := c.BodyParser(userL); err != nil {
		return err
	}
	// 这里登录可以分为两种情况
	// 第一种是使用phoneNumber直接验证
	// 不需要使用密码
	// 或者使用用户Id加密码登录
	if len(userL.Password) == 0 && len(userL.PhoneNumber) > 0 {
		// 将字符串格式的手机号转换成int类型
		phone_num, err := strconv.Atoi(userL.PhoneNumber)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"Message": "请输入正确的手机号",
			})
		}
		// 如果转换成功，说明格式正确
		// 但这不能说明可以使用
		// 这里需要发送验证码
		if err := utils.SendCode(phone_num); err != nil {
			return c.Status(http.StatusOK).JSON(fiber.Map{
				"Message": "输入的手机号有误",
			})
		}
		if ok := utils.ParseCode(userL.Code, phone_num); ok {
			return c.Status(http.StatusOK).JSON(fiber.Map{
				"Message": "登录成功",
			})
		}
	}
	// 第二种情况
	// 使用identity来登录
	if userL.Identity != 0 && len(userL.Password) > 0 {
		// 这种情况下拿着identity和password去数据库里查询下数据
		if result := dao.GetUserByIdentityAndPasswd(userL.Identity, userL.Password); result == 0 {
			return c.Status(http.StatusOK).JSON(fiber.Map{
				"Message": "identity或密码不正确",
			})
		}
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"Message": "登录成功",
	})
}

// 修改自身的信息
func AuthModInfo(c *fiber.Ctx) error {
	return nil
}

// 显示好友
func AuthListFrident(c *fiber.Ctx) error {
	return nil
}

// 添加好友
func AuthAddFriends(c *fiber.Ctx) error {
	return nil
}
