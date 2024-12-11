package response

const (
	ErrCodeSuccess       = 20001 // Success
	ErrCodeParramInvalid = 20003 // Email is invalid
	ErrInvalidToken      = 20004  // Invalid token


)

// message represents
var msg = map[int]string{
	ErrCodeSuccess:       "Success",
	ErrCodeParramInvalid: "Email is invalid",
	ErrInvalidToken:      "Invalid token",
}
