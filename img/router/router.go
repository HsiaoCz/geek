package router

import (
	"log"
	"net/http"
	"time"

	"github.com/HsiaoCz/geek/img/api"
)

func RegRouter(addr string) {
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

	srv := http.Server{
		Handler:      nil,
		Addr:         addr,
		ReadTimeout:  1500 * time.Millisecond,
		WriteTimeout: 1500 * time.Millisecond,
	}
	log.Fatal(srv.ListenAndServe())
}
