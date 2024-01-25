package data

import "gorm.io/gorm"

type UserRepo interface {
	CreateUser()
	DeleteUser()
	ModefyUser()
	GetUserByID(id int64)
}

type UserCase struct {
	db *gorm.DB
}

func NewUserCase() UserRepo {
	return &UserCase{
		db: db,
	}
}

func (u *UserCase) CreateUser()          {}
func (u *UserCase) DeleteUser()          {}
func (u *UserCase) ModefyUser()          {}
func (u *UserCase) GetUserByID(id int64) {}
