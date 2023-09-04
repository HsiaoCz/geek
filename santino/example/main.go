package main

import (
	"math/rand"
	"net/http"
	"strconv"
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

type Book struct {
	Identity int    `json:"identity"`
	BookName string `json:"book_name"`
	Auther   string `json:"auther"`
	Title    string `json:"title"`
	BookType string `json:"book_type"`
	Summery  string `json:"summery"`
}

func GenIdentity() int {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	return random.Intn(1000000)
}

func main() {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	r.POST("/user/register", HandleUserRegister)
	r.GET("/user/:id", HandleGetUserByID)
	r.GET("/book", HandleGetBookByQuery)
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
	user := User{
		Identity: GenIdentity(),
		Username: userR.Username,
		Password: userR.Password,
		Content:  "Hello Hello My man",
		Article:  []string{"斗破苍穹", "我欲封天"},
	}
	c.JSON(http.StatusOK, gin.H{
		"Message": "注册成功",
		"Data":    user,
	})
}

// 这里看一些gin的动态路由
// 这在标准库是实现不了的
// 为什么呢？

// 我们在注册路由的时候，是没法写成这种形式的
// 所以在路由匹配的时候，写成/user/12333
// 也匹配不到我们注册的HandleFunc

// 那么这里的问题就是
// 如何匹配到动态路由呢？
// 如何匹配带有统配符的路由呢？
// 如何简单的实现RESTful API呢？
// 以及中间件等问题，这些问题就是我们实现自己的框架的意义
// 或者说，使用web框架可以帮我们更友好的去实现一些功能

func HandleGetUserByID(c *gin.Context) {
	id := c.Param("id")
	userId, err := strconv.Atoi(id)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}
	user := User{
		Identity: userId,
		Username: "zhangsan",
		Password: "122334",
		Content:  "hello my man",
		Article:  []string{"斗破苍穹", "逆天邪神"},
	}
	c.JSON(http.StatusOK, gin.H{
		"Message": "查询成功",
		"Data":    user,
	})
}

// 这里试一下获取url上的参数
func HandleGetBookByQuery(c *gin.Context) {
	id := c.Query("id")
	bookId, err := strconv.Atoi(id)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}
	book := Book{
		Identity: bookId,
		BookName: "斗破苍穹",
		Auther:   "唐家三少",
		Title:    "中原五白的书",
		BookType: "小白文",
		Summery:  "少年的逆袭之路",
	}
	c.JSON(http.StatusOK, gin.H{
		"Message": "查询成功",
		"Data":    book,
	})
}
