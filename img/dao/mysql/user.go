package mysql

// CreateUser 在数据库创建用户信息

func CreateUser(username string, password string, phoneNumber string) error {}

// GetUserByUsernameAndPhoneNumber 创建用户之前先在数据库里查询一下，避免产生重复数据

func GetUserByUsernameAndPhoneNumber(username string, phoneNumber string) int {}
