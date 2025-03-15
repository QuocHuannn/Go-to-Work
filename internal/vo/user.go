package vo

type UserRegistratorRequest struct {
	Email   string `json:"email" binding:"required,email"`
	Purpose string `json:"purpose" binding:"required"` // TEST_USER, TRADER, ADMIN, etc
}

type UserInfoResponse struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	FullName string `json:"full_name"`
	Status   int    `json:"status"`
}
