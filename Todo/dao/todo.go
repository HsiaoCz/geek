package dao

import (
	"github.com/HsiaoCz/geek/Todo/model"
	"github.com/HsiaoCz/geek/Todo/utils"
)

// 获取todolist
func GetToList(userID int64) []model.TodoList {
	todolist := []model.TodoList{}
	db.Model(&model.TodoList{}).Where("user_id = ?", userID).Scan(todolist)
	return todolist
}

// 添加todo

func AddTodoList(userid int64, content string) error {
	todolist := model.TodoList{
		UserID:   userid,
		Content:  content,
		Identity: utils.GenIdentity(),
	}
	result := db.Create(&todolist)
	return result.Error
}
