package repo

// type UserRepo struct{}

// func NewuserRepo() *UserRepo {
// 	return &UserRepo{}
// }

// // user repo u
// func (ur *UserRepo) GetInfoUser() string {
// 	return "Golang"
// }

// INTERFACE_VERSION
type IUserRepo interface {
	GetUserByEmail(email string) bool
}

type UserRepository struct {
}

// GetUserByEmail implement IUserRepositeory
func (ur *UserRepository) GetUserByEmail(email string) bool {
	return true
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}
