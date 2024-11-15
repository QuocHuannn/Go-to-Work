package service

import "github.com/QuocHuannn/Go-to-Work/internal/repo"

type UserService struct {
	UserRepo *repo.UserRepo
}

func NewuserService() *UserService {
	return &UserService{
		UserRepo: repo.NewuserRepo(),
	}
}

// user repo u
func (us *UserService) GetInfoUser() string {
	return us.UserRepo.GetInfoUser()
}
