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

// 修改todo
func ModTodoList(userid int64, identity int64, content string) error {
	tx := db.Model(&model.TodoList{}).Where("user_id = ? AND identity = ?", userid, identity).Update("content", content)
	return tx.Error
}

// 删除todo
func DeleteTodoList(identity int64, userID int64) error {
	todo := model.TodoList{}
	tx := db.Where("user_id = ? AND identity = ?", userID, identity).Delete(&todo)
	return tx.Error
}

// 获取todo
func GetTodoList(userID int64, ideneity int64) error {
	todo := model.TodoList{}
	tx := db.Model(&todo).Where("user_id = ? AND identity = ?", userID, ideneity).Scan(&todo)
	return tx.Error
}
