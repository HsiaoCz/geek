package controller

// 用户注册时的

type UserR struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	RePasswrod string `json:"re_password"`
	Email      string `json:"email"`
}
