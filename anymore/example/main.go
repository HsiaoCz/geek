package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/HsiaoCz/geek/anymore"
)

type UserR struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	RePassword string `json:"re_password"`
}

func main() {
	r := anymore.App()
	r.POST("/user/register", UserRegister)
	r.GET("/user/id", GetUserById)
	r.Listen("127.0.0.1:9091")
}

func UserRegister(c *anymore.Context) {
	userR := new(UserR)
	if err := json.NewDecoder(c.R.Body).Decode(userR); err != nil {
		return
	}
	c.JSON(http.StatusOK, anymore.H{
		"Message": "注册成功",
	})
}

func GetUserById(c *anymore.Context) {
	id := c.Query("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, anymore.H{
			"Message": "请输入正确ID",
		})
		return
	}
	c.JSON(http.StatusOK, anymore.H{
		"Message": "获取成功",
		"Data":    userID,
	})
}
