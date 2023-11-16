package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./index.html", "./login.html", "message.html")
	r.GET("/app/v1/auther", HandleUser)
	r.POST("/app/v1/auther/login", HandleUserLogin)
	log.Fatal(r.Run("127.0.0.1:3002"))
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Content  string `json:"content"`
}

func HandleUser(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func HandleUserLogin(c *gin.Context) {
	username := c.Request.FormValue("username")
	password := c.Request.FormValue("password")
	rePassword := c.Request.FormValue("re_password")
	if len(username) == 0 || len(password) == 0 || len(rePassword) == 0 {
		c.HTML(http.StatusBadRequest, "message.html", "请输入用户名或密码")
		return
	}
	if password != rePassword {
		c.HTML(http.StatusBadRequest, "message.html", "请确认密码")
		return
	}
	user := User{
		Username: username,
		Password: password,
		Content:  "我们意念合一",
	}
	c.HTML(http.StatusOK, "index.html", &user)
}
