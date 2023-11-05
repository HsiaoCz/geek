package controller

type UserR struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	RePassword  string `json:"re_password"`
	PhoneNumber string `json:"phone_number"`
}

type UserL struct {
	// 有两种方式
	// 第一种是直接通过id 登录
	// 第二种是通过电话登录
	Identity    int    `json:"identity"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}
