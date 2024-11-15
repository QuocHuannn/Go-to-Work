package controller

import (
	"net/http"

	"github.com/QuocHuannn/Go-to-Work/internal/service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		UserService: service.NewuserService(),
	}
}

// uc user controller
// us user service
// controller -> service -> repo -> models -> database
func (uc *UserController) GetUserByID(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": uc.UserService.GetInfoUser(),
		"users":   []string{"user1", "user2"},
	})
}
