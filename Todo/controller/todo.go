package controller

import (
	"log"
	"net/http"

	"github.com/HsiaoCz/geek/Todo/dao"
	"github.com/gin-gonic/gin"
)

func AddTodoList(c *gin.Context) {
	todo := new(Todo)
	c.Bind(todo)
	if len(todo.Content) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"Message": "请输入待做事项的内容",
		})
		return
	}
	if result := dao.AuthGetUserID(todo.UserId); result == 0 {
		c.JSON(http.StatusOK, gin.H{
			"Message": "当前用户id有误",
		})
		return
	}
	if err := dao.AddTodoList(todo.UserId, todo.Content); err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK,gin.H{
		"Message":"创建todo成功",
	})
}
func ModTodoList(c *gin.Context)      {}
func DeleteTodoList(c *gin.Context)   {}
func GetTodoList(c *gin.Context)      {}
func CompleteTodoList(c *gin.Context) {}
func CleanTodoList(c *gin.Context)    {}
