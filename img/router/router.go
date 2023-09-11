package router

import (
	"net/http"

	"github.com/HsiaoCz/geek/img/api"
)

func RegRouter() {
	// 用户注册
	http.HandleFunc("/img/user/register", api.HandleUserRegister)
	// 用户登录
	http.HandleFunc("/img/user/login", api.HandleUserLogin)
	// 添加好友
	http.HandleFunc("/img/user/addf", api.HandleUserAddFridents)
	// 添加分组
	http.HandleFunc("/img/user/groupf", api.HandleUserGroupf)
	// 删除好友
	http.HandleFunc("/img/user/delf", api.HandleUserDeleteF)
	// 移动好友到分组
	http.HandleFunc("/img/user/movef", api.HandleUserMoveF)
	// 修改好友的昵称
	http.HandleFunc("/img/user/modif", api.HandleUserModif)
}
