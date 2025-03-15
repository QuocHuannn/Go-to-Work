package controller

import (
	"net/http"

	"github.com/QuocHuannn/Go-to-Work/internal/service"
	"github.com/QuocHuannn/Go-to-Work/internal/vo"
	"github.com/QuocHuannn/Go-to-Work/pkg/response"
	"github.com/gin-gonic/gin"
)

// import (
// 	"net/http"

// 	"github.com/QuocHuannn/Go-to-Work/internal/service"
// 	"github.com/gin-gonic/gin"
// )

type OTPRequest struct {
	Email string `json:"email" binding:"required,email"`
	OTP   string `json:"otp" binding:"required"`
}

type UserController struct {
	UserService service.IUserService
}

func NewUserController(
	userService service.IUserService,
) *UserController {
	return &UserController{
		UserService: userService,
	}
}

// Home Page
func (uc *UserController) HomePage(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", gin.H{
		"title": "CRM System - Home",
	})
}

// API Endpoints (JSON)

func (uc *UserController) Register(c *gin.Context) {
	var req vo.UserRegistratorRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ResponseError(c, response.ErrCodeParramInvalid, err.Error())
		return
	}

	result := uc.UserService.Register(c.Request.Context(), req.Email, req.Purpose)
	response.ResponseSuccess(c, result)
}

func (uc *UserController) VerifyOTP(c *gin.Context) {
	var req OTPRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ResponseError(c, response.ErrCodeParramInvalid, err.Error())
		return
	}

	result := uc.UserService.VerifyOTP(c.Request.Context(), req.Email, req.OTP)
	response.ResponseSuccess(c, result)
}

func (uc *UserController) GetUserInfo(c *gin.Context) {
	email := c.Query("email")
	if email == "" {
		response.ResponseError(c, response.ErrCodeParramInvalid, "Email is required")
		return
	}

	user, err := uc.UserService.GetUserByEmail(c.Request.Context(), email)
	if err != nil {
		response.ResponseError(c, response.ErrCodeParramInvalid, err.Error())
		return
	}

	response.ResponseSuccess(c, user)
}

// HTML Templates

func (uc *UserController) RegisterPage(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{
		"title": "User Registration",
	})
}

func (uc *UserController) VerifyOTPPage(c *gin.Context) {
	email := c.Query("email")
	c.HTML(http.StatusOK, "verify-otp.html", gin.H{
		"title": "Verify OTP",
		"email": email,
	})
}

func (uc *UserController) ProfilePage(c *gin.Context) {
	email := c.Query("email")
	if email == "" {
		c.Redirect(http.StatusFound, "/register")
		return
	}

	user, err := uc.UserService.GetUserByEmail(c.Request.Context(), email)
	if err != nil {
		c.HTML(http.StatusOK, "error.html", gin.H{
			"title": "Error",
			"error": "User not found",
		})
		return
	}

	c.HTML(http.StatusOK, "profile.html", gin.H{
		"title": "User Profile",
		"user":  user,
	})
}

// Form handler - Process register form submission
func (uc *UserController) ProcessRegister(c *gin.Context) {
	email := c.PostForm("email")
	purpose := c.PostForm("purpose")

	if email == "" || purpose == "" {
		c.HTML(http.StatusBadRequest, "register.html", gin.H{
			"title": "User Registration",
			"error": "Email and purpose are required",
		})
		return
	}

	result := uc.UserService.Register(c.Request.Context(), email, purpose)
	if result != response.ErrCodeSuccess {
		var errorMsg string
		if result == response.ErrCodeUserHasExist {
			errorMsg = "Email already registered"
		} else {
			errorMsg = "Registration failed"
		}

		c.HTML(http.StatusBadRequest, "register.html", gin.H{
			"title": "User Registration",
			"error": errorMsg,
		})
		return
	}

	// Redirect to OTP verification page
	c.Redirect(http.StatusFound, "/verify-otp?email="+email)
}

// Form handler - Process OTP verification form submission
func (uc *UserController) ProcessVerifyOTP(c *gin.Context) {
	email := c.PostForm("email")
	otp := c.PostForm("otp")

	if email == "" || otp == "" {
		c.HTML(http.StatusBadRequest, "verify-otp.html", gin.H{
			"title": "Verify OTP",
			"email": email,
			"error": "Email and OTP are required",
		})
		return
	}

	result := uc.UserService.VerifyOTP(c.Request.Context(), email, otp)
	if result != response.ErrCodeSuccess {
		c.HTML(http.StatusBadRequest, "verify-otp.html", gin.H{
			"title": "Verify OTP",
			"email": email,
			"error": "Invalid OTP or verification failed",
		})
		return
	}

	// Redirect to profile page
	c.Redirect(http.StatusFound, "/profile?email="+email)
}

// // uc user controller
// // us user service
// // controller -> service -> repo -> models -> database

// func (uc *UserController) GetUserByID(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": uc.UserService.GetInfoUser(),
// 		"users":   []string{"user1", "user2"},
// 	})
// }
