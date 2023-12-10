package av1

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/HsiaoCz/geek/what/api/define"
	"github.com/HsiaoCz/geek/what/pb"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func UserSinup(c *fiber.Ctx) error {
	userS := new(define.UserS)
	if err := c.BodyParser(userS); err != nil {
		return err
	}
	if len(userS.Username) == 0 || len(userS.Password) == 0 || len(userS.RePassword) == 0 || len(userS.Email) == 0 {
		return c.Status(http.StatusOK).JSON(fiber.Map{
			"Message": "请填写用户名或密码或邮箱",
			"Code":    10002,
		})
	}
	if userS.Password != userS.RePassword {
		return c.Status(http.StatusOK).JSON(fiber.Map{
			"Message": "请检查用户名或密码",
			"code":    10003,
		})
	}
	if ok := strings.Contains(userS.Email, "@"); !ok {
		return c.Status(http.StatusOK).JSON(fiber.Map{
			"Message": "请输入有效邮箱",
			"Code":    10004,
		})
	}
	conn, err := grpc.Dial(define.UserServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	defer conn.Close()
	client := pb.NewWhatClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	result, err := client.UserSinup(ctx, &pb.SinupRequest{Username: userS.Username, Password: userS.Password, Repassword: userS.RePassword, Email: userS.Email})
	if err != nil {
		return c.Status(http.StatusOK).JSON(fiber.Map{
			"Message": result.GetMsg(),
			"Code":    result.GetCode(),
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"Message": result.GetMsg(),
		"Code":    result.GetCode(),
	})
}

func UserLogin(c *fiber.Ctx) error {
	userL := new(define.UserL)
	if err := c.BodyParser(userL); err != nil {
		return err
	}
	if len(userL.Username) == 0 || len(userL.Password) == 0 || len(userL.Email) == 0 {
		return c.Status(http.StatusOK).JSON(fiber.Map{
			"Message": "用户名或密码或邮箱不能为空",
			"Code":    10002,
		})
	}
	if ok := strings.Contains(userL.Email, "@"); !ok {
		return c.Status(http.StatusOK).JSON(fiber.Map{
			"Message": "请输入有效邮箱",
			"Code":    10003,
		})
	}
	conn, err := grpc.Dial(define.UserServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	defer conn.Close()
	client := pb.NewWhatClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	result, err := client.UserLogin(ctx, &pb.LoginRequest{Username: userL.Username, Password: userL.Password, Email: userL.Email})
	if err != nil {
		return c.Status(http.StatusOK).JSON(fiber.Map{
			"Message": result.GetMsg(),
			"Code":    result.GetCode(),
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"Message": result.GetMsg(),
		"Code":    result.GetCode(),
	})
}
