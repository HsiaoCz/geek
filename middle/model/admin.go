package model

// 管理员的结构体
// 管理员可以审核文章
// 但这里的问题是，这个管理员的类应该有哪些字段呢?
// 这里暂时先写这些
type Admin struct {
	Username string `json:"username"`
	Identity int    `json:"identity"`
	Password string `json:"password"`
	Count    int    `json:"count"`
}
