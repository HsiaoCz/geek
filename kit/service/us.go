package service

type IUserService interface {
	GetUsername(userID int) string
}

type UserService struct{}

func (u *UserService) GetUsername(userID int) string {
	if userID == 100001 {
		return "bob"
	}
	return "alex"
}
