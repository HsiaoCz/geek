package sso

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// 单机登录
// 先启动一个业务服务器，在业务服务器上实现一个单机登录的过程
func TestBizServer(t *testing.T) {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	r.GET("/profile", Profile)
	r.POST("/login", Login)
	t.Fatal(r.Run(":9091"))
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Age      string `json:"age"`
}

func Profile(c *gin.Context) {
	proUser := &User{
		Username: "zhangsan",
		Age:      "liss",
	}
	c.JSON(http.StatusOK, proUser)
}

// 使用一个map来存储session
var sessions map[string]interface{}

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if username != "zhangsan" || password != "lisi" {
		c.JSON(http.StatusOK, gin.H{
			"Message": "用户名或密码错误",
		})
	}
	userID := uuid.New().String()
	c.SetCookie("token", userID, 3600, "/", "localhost", false, true)
	sessions[userID] = username
	// 这里的关键在于
	// 这个session信息需要保存到一个位置
	// 另外要给他设置过期时间
	// 最好的地方就是redis
	// 获取cookie
	usID, err := c.Cookie("token")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"Message": "请登录",
		})
		return
	}

	usname, ok := sessions[usID]
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"message": "请登录",
		})
		return
	}
	if usname == username {
		c.Next()
	}
}
