package random

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	// Initialize random seed
	rand.Seed(time.Now().UnixNano())
}

// GenerateSixDigitsOTP generates a random 6-digit OTP
func GenerateSixDigitsOTP() string {
	otp := rand.Intn(900000) + 100000 // Generates number between 100000 and 999999
	return fmt.Sprintf("%06d", otp)
}

func GenerateRandomPassword() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+"
	const length = 12

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
