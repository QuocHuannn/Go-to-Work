package service

import (
	"github.com/QuocHuannn/Go-to-Work/internal/repo"
	"github.com/QuocHuannn/Go-to-Work/pkg/response"
)

// type UserService struct {
// 	UserRepo *repo.UserRepo
// }

// func NewuserService() *UserService {
// 	return &UserService{
// 		UserRepo: repo.NewuserRepo(),
// 	}
// }

// // user repo u
// func (us *UserService) GetInfoUser() string {
// 	return us.UserRepo.GetInfoUser()
// }

// INTERFACE VERSION
type IUserService interface {
	Register(email string, purpose string) int
}
type UserService struct {
	UserRepo *repo.UserRepository
}
func NewUserService(
	userRepo *repo.UserRepository,
) IUserService {
	return &UserService{
		UserRepo: userRepo,
	}
}


// Register implements IUserService
func (us *UserService) Register(email string, purpose string) int {
	//1. Check if email is already registered
	if us.UserRepo.GetUserByEmail(email) {
		return response.ErrCodeUserHasExist
	}
	return response.ErrCodeSuccess
}
