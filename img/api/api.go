package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/HsiaoCz/geek/img/dao/mysql"
)

// HandleUserRegister 用户注册
func HandleUserRegister(w http.ResponseWriter, r *http.Request) {
	userR := new(UserR)
	err := json.NewDecoder(r.Body).Decode(userR)
	if err != nil {
		ResponseJSON(w, http.StatusOK, H{
			"Msg": err,
		})
		return
	}
	if userR.Password != userR.RePassword {
		ResponseJSON(w, http.StatusOK, H{
			"Msg": "请检查用户名或密码",
		})
		return
	}
	if count := mysql.GetUserByUsernameAndPhoneNumber(userR.Username, userR.PhoneNumber); count != 0 {
		ResponseJSON(w, http.StatusOK, H{
			"Msg": "当前用户已经注册",
		})
		return
	}
	if err := mysql.CreateUser(userR.Username, userR.Password, userR.PhoneNumber); err != nil {
		log.Println(err)
		return
	}
	ResponseJSON(w, http.StatusOK, H{
		"Msg": "注册成功",
	})
}

// HandleUserLogin 用户登录
func HandleUserLogin(w http.ResponseWriter, r *http.Request) {
	userL := new(UserL)
	if err := json.NewDecoder(r.Body).Decode(userL); err != nil {
		log.Println(err)
		return
	}
	if err := mysql.GetUserByPhoneNuberAndPassword(userL.PhoneNumber, userL.Password); err != nil {
		ResponseJSON(w, http.StatusOK, H{
			"Msg": "用户名或密码错误",
		})
		return
	}
	ResponseJSON(w, http.StatusOK, H{
		"Msg": "登录成功",
	})
}

// HandleUserAddFridents 添加好友
func HandleUserAddFridents(w http.ResponseWriter, r *http.Request) {}

// HandleUserGroup 添加分组信息
func HandleUserGroupf(w http.ResponseWriter, r *http.Request) {}

// HandleUserDeleteF 删除好友
func HandleUserDeleteF(w http.ResponseWriter, r *http.Request) {}

// HandleUserMoveF 移动好友到某个分组
func HandleUserMoveF(w http.ResponseWriter, r *http.Request) {}

// HandleUserModif 给好友添加备注
func HandleUserModif(w http.ResponseWriter, r *http.Request) {}
