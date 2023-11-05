package dao

import (
	"github.com/HsiaoCz/geek/Tenz/model"
	"github.com/HsiaoCz/geek/Tenz/utils"
)

// 根据电话号判断用户表是否有用户使用此电话号注册过
func AuthTestPhoneNumber(phoneNumber string) int64 {
	var user model.User
	result := db.Model(&user).Where("phone_number = ?", phoneNumber).Find(&user)
	return result.RowsAffected
}

// 创建用户
func AuthCreate(username string, password string, phoneNumber string) error {
	user := model.User{
		Username:    username,
		Passwrod:    utils.SetMd5Password(password),
		PhoneNumber: phoneNumber,
		Level:       0,
		Identity:    utils.GenIdentity(),
		Job:         "",
		Signature:   "",
		Content:     "",
		Birthday:    "",
	}
	tx := db.Create(&user)
	return tx.Error
}
