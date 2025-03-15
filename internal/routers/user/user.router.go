package user

import (
	"net/http"

	"github.com/QuocHuannn/Go-to-Work/global"
	"github.com/QuocHuannn/Go-to-Work/internal/wire"
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (pr *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userController, _ := wire.InitUserRouterHandler(global.Mdb)

	// Home page
	Router.GET("/", userController.HomePage)

	// HTML pages
	Router.GET("/register", userController.RegisterPage)
	Router.GET("/verify-otp", userController.VerifyOTPPage)
	Router.GET("/profile", userController.ProfilePage)

	// Form submission handlers
	Router.POST("/register", userController.ProcessRegister)
	Router.POST("/verify-otp", userController.ProcessVerifyOTP)

	// API endpoints
	userRouterAPI := Router.Group("/api")
	{
		userRouterAPI.POST("/user/register", userController.Register)
		userRouterAPI.POST("/user/otp", userController.VerifyOTP)
		userRouterAPI.GET("/user/get_info", userController.GetUserInfo)
	}

	// Admin API
	adminRouter := Router.Group("/admin")
	// adminRouter.Use(middleware.JWTAuth())
	{
		adminRouter.GET("/user/active_user", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"status":  "ok",
				"message": "Active users endpoint",
				"data": []map[string]interface{}{
					{"id": 1, "username": "user1", "active": true},
					{"id": 2, "username": "user2", "active": true},
				},
			})
		})
	}
}
