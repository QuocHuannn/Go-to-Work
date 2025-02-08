package response

const (
	ErrCodeSuccess       = 20001 // Success
	ErrCodeParramInvalid = 20003 // Email is invalid
	ErrInvalidToken      = 30001 // Invalid token
	// Register Code
	ErrCodeUserHasExist = 5001 // User already exists

)

// message represents
var msg = map[int]string{
	ErrCodeSuccess:       "Success",
	ErrCodeParramInvalid: "Email is invalid",
	ErrInvalidToken:      "Invalid token",
	ErrCodeUserHasExist:  "User already exists",
}
