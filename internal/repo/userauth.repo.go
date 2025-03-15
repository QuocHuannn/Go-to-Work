package repo

import (
	"context"
	"fmt"
	"time"

	"github.com/QuocHuannn/Go-to-Work/global"
	"gorm.io/gorm"
)

type IUserAuthRepository interface {
	AddOTP(ctx context.Context, email string, otp string, expirationTime int64) error
	VerifyOTP(ctx context.Context, email string, otp string) (bool, error)
}

type userAuthRepository struct {
	db *gorm.DB
}

func (u *userAuthRepository) AddOTP(ctx context.Context, email string, otp string, expirationTime int64) error {
	key := fmt.Sprintf("user:%s:otp", email)
	return global.Rdb.Set(ctx, key, otp, time.Duration(expirationTime)*time.Second).Err()
}

func (u *userAuthRepository) VerifyOTP(ctx context.Context, email string, otp string) (bool, error) {
	key := fmt.Sprintf("user:%s:otp", email)
	storedOTP, err := global.Rdb.Get(ctx, key).Result()
	if err != nil {
		return false, err
	}
	return storedOTP == otp, nil
}

func NewUserAuthRepository(db *gorm.DB) IUserAuthRepository {
	return &userAuthRepository{
		db: db,
	}
}
