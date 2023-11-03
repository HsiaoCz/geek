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

// 根据用户名和邮箱和密码去获取用户的信息
func AuthGetUserInfoByUsernameAndPasswd(username string, password string) *model.User {
	user := new(model.User)
	db.Model(user).Where("username = ? AND password = ?", username, utils.SetMd5Password(password)).Scan(user)
	return user
}

// 查找当前用户id是否存在
func AuthGetUserID(userID int64) int64 {
	user := new(model.User)
	result := db.Where("user_id = ? ", userID).Find(user)
	return result.RowsAffected
}
