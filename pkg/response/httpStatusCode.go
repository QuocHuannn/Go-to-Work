package response

const (
	ErrCodeSuccess       = 20001 // Success
	ErrCodeParramInvalid = 20003 // Email is invalid

	ErrInvalidToken     = 30001 // Invalid token
	ErrInvalidOTP       = 30002 // Invalid OTP
	ErrSendEmailOtpFail = 30003 // Send email OTP fail
	// Register Code
	ErrCodeUserHasExist = 5001 // User already exists

)

// message represents
var msg = map[int]string{
	ErrCodeSuccess:       "Success",
	ErrCodeParramInvalid: "Email is invalid",
	ErrInvalidToken:      "Invalid token",
	ErrInvalidOTP:        "Invalid OTP",
	ErrSendEmailOtpFail:  "Send email OTP fail",
	ErrCodeUserHasExist:  "User already exists",
}
