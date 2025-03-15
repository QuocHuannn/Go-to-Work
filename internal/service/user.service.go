package service

import (
	"context"
	"fmt"
	"time"

	"github.com/QuocHuannn/Go-to-Work/internal/config"
	"github.com/QuocHuannn/Go-to-Work/internal/model"
	"github.com/QuocHuannn/Go-to-Work/internal/repo"
	"github.com/QuocHuannn/Go-to-Work/internal/utils/crypto"
	"github.com/QuocHuannn/Go-to-Work/internal/utils/random"
	"github.com/QuocHuannn/Go-to-Work/internal/utils/sendto"
	"github.com/QuocHuannn/Go-to-Work/internal/vo"
	"github.com/QuocHuannn/Go-to-Work/pkg/response"
)

type IUserService interface {
	Register(ctx context.Context, email string, purpose string) int
	VerifyOTP(ctx context.Context, email string, otp string) int
	GetUserByEmail(ctx context.Context, email string) (*vo.UserInfoResponse, error)
}

type UserService struct {
	UserRepo     repo.IUserRepository
	UserAuthRepo repo.IUserAuthRepository
}

func NewUserService(
	userRepo repo.IUserRepository,
	userAuthRepo repo.IUserAuthRepository,
) IUserService {
	return &UserService{
		UserRepo:     userRepo,
		UserAuthRepo: userAuthRepo,
	}
}

// Register implements IUserService
func (us *UserService) Register(ctx context.Context, email string, purpose string) int {
	//0. Hash email
	hashEmail := crypto.HashEmail(email)
	fmt.Printf("Hash email is : %s\n", hashEmail)

	//1. Check if email is already registered
	exists, _ := us.UserRepo.GetUserByEmail(email)
	if exists != nil {
		return response.ErrCodeUserHasExist
	}

	//2. Generate new OTP
	var otp string
	if purpose == "TEST_USER" {
		otp = "123456" // Use predefined OTP for testing
	} else {
		otp = random.GenerateSixDigitsOTP() // Generate random OTP for real users
	}

	fmt.Println("OTP is : ", otp)

	//3. save OTP in Redis with expiration time
	err := us.UserAuthRepo.AddOTP(ctx, hashEmail, otp, int64(10*time.Minute))
	if err != nil {
		return response.ErrInvalidOTP
	}

	//4. send OTP to email
	recipients := []string{email}
	err = sendto.SendTextEmailOTP(recipients, config.Cfg.SMTP.FromEmail, otp)
	if err != nil {
		return response.ErrSendEmailOtpFail
	}

	return response.ErrCodeSuccess
}

// VerifyOTP implements IUserService
func (us *UserService) VerifyOTP(ctx context.Context, email string, otp string) int {
	// Hash email
	hashEmail := crypto.HashEmail(email)

	// Verify OTP from Redis
	isValid, err := us.UserAuthRepo.VerifyOTP(ctx, hashEmail, otp)
	if err != nil {
		return response.ErrInvalidOTP
	}

	if !isValid {
		return response.ErrInvalidOTP
	}

	// Create user in database with status active (1)
	// (This would be a separate API in a typical flow, but for demo
	// we're creating the user after OTP verification)

	// Create random password
	password := random.GenerateRandomPassword()
	hashedPassword := crypto.HashPassword(password)

	user := &model.User{
		Email:    email,
		Password: hashedPassword,
		FullName: email, // Use email as fullname for now
		Status:   1,     // Set to active
	}

	err = us.UserRepo.CreateUser(user)
	if err != nil {
		return response.ErrCodeParramInvalid
	}

	return response.ErrCodeSuccess
}

// GetUserByEmail returns user info by email
func (us *UserService) GetUserByEmail(ctx context.Context, email string) (*vo.UserInfoResponse, error) {
	user, err := us.UserRepo.GetUserByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	return &vo.UserInfoResponse{
		ID:       user.ID,
		Email:    user.Email,
		FullName: user.FullName,
		Status:   user.Status,
	}, nil
}
