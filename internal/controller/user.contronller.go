package controller

import (
	"github.com/QuocHuannn/Go-to-Work/internal/service"
	"github.com/QuocHuannn/Go-to-Work/pkg/response"
	"github.com/gin-gonic/gin"
)

// import (
// 	"net/http"

// 	"github.com/QuocHuannn/Go-to-Work/internal/service"
// 	"github.com/gin-gonic/gin"
// )

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

func (uc *UserController) Register(c *gin.Context) {
	result := uc.UserService.Register("", "")
	response.ResponseSuccess(c, result)
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
