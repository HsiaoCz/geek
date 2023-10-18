package model

type User struct {
	Identity string `json:"identity"`
	Username string `json:"username"`
	Password string `json:"password"`
	Content  string `json:"content"`
}
