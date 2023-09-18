package api

// 用户注册
type UserR struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	RePassword  string `json:"re_password"`
	PhoneNumber string `json:"phone_number"`
}

// 用户登录
type UserL struct {
	Id          int    `json:"id"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
}
