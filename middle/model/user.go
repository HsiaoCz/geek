package model

type User struct {
	Username string `json:"username"`
	Identity int    `json:"identity"`
	Password string `json:"passwrod"`
	Job      string `json:"job"`
	Content  string `json:"content"`
	Email    string `json:"email"`
}
