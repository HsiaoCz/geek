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

// 利用identity和password来数据库里面查询一下记录
func GetUserByIdentityAndPasswd(identity int, password string) int64 {
	user := model.User{}
	tx := db.Where("identity = ? AND password = ?", identity, utils.SetMd5Password(password)).Find(&user)
	return tx.RowsAffected
}
