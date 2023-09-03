package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	Addr = "127.0.0.1:9092"
)

type User struct {
	Identity int      `json:"identity"`
	Username string   `json:"username"`
	Password string   `json:"password"`
	Content  string   `json:"content"`
	Article  []string `json:"articles"`
}

type UserRegister struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	RePassword string `json:"re_password"`
}

func main() {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	r.POST("/user/register", HandleUserRegister)
	srv := http.Server{
		Handler:      r,
		Addr:         Addr,
		WriteTimeout: 1500 * time.Millisecond,
		ReadTimeout:  1500 * time.Millisecond,
	}
	srv.ListenAndServe()
}

func HandleUserRegister(c *gin.Context) {
	var userR UserRegister
	err := c.BindJSON(&userR)
	if err != nil {
		return
	}
	if userR.Password != userR.RePassword {
		c.JSON(http.StatusOK, gin.H{
			"Message": "请输入相同的密码和确认密码",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Message": "注册成功",
	})
}
