package controller

import (
	"log"
	"net/http"
	"strings"

	"github.com/HsiaoCz/geek/Todo/dao"
	"github.com/HsiaoCz/geek/Todo/utils"
	"github.com/gin-gonic/gin"
)

// AuthRegister 用户注册
func AuthRegister(c *gin.Context) {
	userR := new(UserR)
	err := c.Bind(userR)
	if err != nil {
		log.Fatal(err)
	}
	if len(userR.Username) == 0 || len(userR.Password) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"Message": "用户名或密码不能为空",
		})
		return
	}
	if userR.Password != userR.RePassword {
		c.JSON(http.StatusOK, gin.H{
			"Message": "请检查密码",
		})
		return
	}
	if !strings.Contains(userR.Email, "@") {
		c.JSON(http.StatusOK, gin.H{
			"Message": "请输入合法的邮箱",
		})
		return
	}
	if result := dao.AuthGetUserByUsernameAndEmail(userR.Username, userR.Email); result > 0 {
		c.JSON(http.StatusOK, gin.H{
			"Message": "该用户当前已经存在",
		})
		return
	}
	if err := dao.AuthReg(userR.Username, userR.Password, utils.GenIdentity(), userR.Email); err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"Message": "注册成功",
	})
}

// AuthLogin 用户登录
func AuthLogin(c *gin.Context) {
	userL := new(UserL)
	err := c.Bind(userL)
	if err != nil {
		log.Fatal(err)
	}
	if result := dao.AuthGetPasswdAndEmial(userL.Username, userL.Password, userL.Email); result == 0 {
		c.JSON(http.StatusOK, gin.H{
			"Message": "用户名或密码不正确",
		})
		return
	}
	user := dao.AuthGetUserInfoByUsernameAndPasswd(userL.Username, userL.Password)
	token, err := utils.GenToken(user.Identity)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"Message": "登录成功",
		"Token":   token,
	})
}

// AuthGetTodoList 获取用户的待做事项列表
func AuthGetTodoList(c *gin.Context) {
	data, ok := c.Get("userID")
	if !ok {
		log.Println("需要设置请求上下文信息")
		return
	}
	userID := data.(int64)
	todolist := dao.GetToList(int64(userID))
	c.JSON(http.StatusOK, gin.H{
		"Message": "获取待做事项成功",
		"Data":    todolist,
	})
}
