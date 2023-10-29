package model

import "gorm.io/gorm"

type TodoList struct {
	gorm.Model
	Identity int64  `gorm:"column:identity" json:"identity"`
	UserID   int64  `gorm:"column:user_id" json:"user_id"`
	Content  string `gorm:"column:content" json:"content"`
}

func (t TodoList)TableName()string{
	return "todolist"
}