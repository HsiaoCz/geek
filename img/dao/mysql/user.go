package mysql

import (
	"github.com/HsiaoCz/geek/img/model"
	"github.com/HsiaoCz/geek/img/settings"
)

// CreateUser 在数据库创建用户信息

func CreateUser(username string, password string, phoneNumber string) error {
	user := &model.User{
		Username:    username,
		Passwrod:    settings.GenMD5(password),
		Identity:    settings.GenIdentity(),
		PhoneNumber: phoneNumber,
	}
	tx := DB.db.Create(user)
	return tx.Error
}

// GetUserByUsernameAndPhoneNumber 创建用户之前先在数据库里查询一下，避免产生重复数据

func GetUserByUsernameAndPhoneNumber(username string, phoneNumber string) int64 {
	tx := DB.db.Where("username = ? AND phone_number = ?", username, phoneNumber).Take(&model.User{})
	return tx.RowsAffected
}

// GetUserByUsernameAndPassword 用户登录时验证登录信息

func GetUserByPhoneNuberAndPassword(phoneNumber string, password string) error {
	tx := DB.db.Where("phone_number = ? AND password = ?", phoneNumber, settings.GenMD5(password)).Take(&model.User{})
	return tx.Error
}
