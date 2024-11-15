package repo

type UserRepo struct{}

func NewuserRepo() *UserRepo {
	return &UserRepo{}
}

// user repo u
func (ur *UserRepo) GetInfoUser() string {
	return "Golang"
}
