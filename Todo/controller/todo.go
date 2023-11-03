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
	c.JSON(http.StatusOK, gin.H{
		"Message": "创建todo成功",
	})
}
func ModTodoList(c *gin.Context) {
	todo := new(Todo)
	if err := c.Bind(todo); err != nil {
		log.Fatal(err)
	}
	if len(todo.Content) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"Message": "请输入内容",
		})
		return
	}
	if err := dao.ModTodoList(todo.Identity, todo.UserId, todo.Content); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"Message": "更新失败",
			"Error":   err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Message": "更新成功",
	})
}
func DeleteTodoList(c *gin.Context) {
	todo := new(Todo)
	if err := c.Bind(todo); err != nil {
		log.Fatal(err)
	}
	if err := dao.DeleteTodoList(todo.Identity, todo.UserId); err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"Message": "删除成功",
	})
}
func GetTodoList(c *gin.Context) {
	todo := new(Todo)
	if err := c.Bind(todo); err != nil {
		log.Fatal(err)
	}
	
}
func CompleteTodoList(c *gin.Context) {}
func CleanTodoList(c *gin.Context)    {}
