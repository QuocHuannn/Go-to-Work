package user

import (
	"github.com/QuocHuannn/Go-to-Work/internal/wire"
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (pr *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userController, _ := wire.InitUserRouterHandler()
	//WIRE GO
	//DEPENDENCY INJECTION
	// public router
	userRouterPublic := Router.Group("/user")
	{
		userRouterPublic.POST("/register", userController.Register) // register --> yes --> no
		userRouterPublic.POST("/otp")
	}
	//private router
	userRouterPrivate := Router.Group("/user")
	//userRouterPrivate.Use(limiter())
	//userRouterPrivate.Use(Authen())
	//userRouterPrivate.Use(Permission())
	{
		userRouterPrivate.GET("/get_info")
	}
}
