package controller

type UserR struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	RePassword string `json:"re_password"`
	Email      string `json:"email"`
}

type UserL struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type Todo struct {
	UserId   int64  `json:"user_id"`
	Content  string `json:"content"`
	Identity int64  `json:"identity"`
}
