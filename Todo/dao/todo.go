package dao

import "github.com/HsiaoCz/geek/Todo/model"

// 获取todolist
func GetToList(userID int64) []model.TodoList {
	todolist := []model.TodoList{}
	db.Model(&model.TodoList{}).Where("user_id = ?", userID).Scan(todolist)
	return todolist
}
