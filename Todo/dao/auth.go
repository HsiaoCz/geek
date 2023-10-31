package dao

import (
	"github.com/HsiaoCz/geek/Todo/model"
	"github.com/HsiaoCz/geek/Todo/utils"
)

// 创建用户
func AuthReg(username string, password string, identity int64, email string) error {
	user := model.User{
		Username: username,
		Password: utils.SetMd5Password(password),
		Identity: identity,
		Email:    email,
		Avatar:   "",
	}
	result := db.Create(&user)
	return result.Error
}

// 用户登录时的验证
func AuthGetPasswdAndEmial(username string, password string, email string) int64 {
	user := new(model.User)
	result := db.Where("username = ? AND password = ? AND email = ?", username, utils.SetMd5Password(password), email).Find(user)
	return result.RowsAffected
}

// 根据用户名和邮箱判断当前数据库中是否存在该用户
func AuthGetUserByUsernameAndEmail(username string, email string) int64 {
	user := new(model.User)
	result := db.Where("username = ? AND email = ?", username, email).Find(user)
	return result.RowsAffected
}