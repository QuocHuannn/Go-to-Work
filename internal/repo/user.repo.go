package repo

import (
	"github.com/QuocHuannn/Go-to-Work/internal/model"
	"gorm.io/gorm"
)

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

type IUserRepository interface {
	GetUserByEmail(email string) (*model.User, error)
	CreateUser(user *model.User) error
	UpdateUser(user *model.User) error
}

type UserRepository struct {
	db *gorm.DB
}

// GetUserByEmail implement IUserRepositeory
func (ur *UserRepository) GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	err := ur.db.WithContext(ctx).Table(TableNameGoCrmUser).Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *UserRepository) CreateUser(user *model.User) error {
	return ur.db.WithContext(ctx).Table(TableNameGoCrmUser).Create(user).Error
}

func (ur *UserRepository) UpdateUser(user *model.User) error {
	return ur.db.WithContext(ctx).Table(TableNameGoCrmUser).Save(user).Error
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{db: db}
}
