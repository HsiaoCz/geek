package api

import "net/http"

// HandleUserRegister 用户注册
func HandleUserRegister(w http.ResponseWriter, r *http.Request) {}

// HandleUserLogin 用户登录
func HandleUserLogin(w http.ResponseWriter, r *http.Request) {}

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
