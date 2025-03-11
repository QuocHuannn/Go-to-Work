package manager

import (
	"github.com/QuocHuannn/Go-to-Work/internal/controller"
	"github.com/QuocHuannn/Go-to-Work/internal/repo"
	"github.com/QuocHuannn/Go-to-Work/internal/service"
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (pr *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	//public router
	// this is non-dependency router
	ur := repo.NewUserRepository()
	us := service.NewUserService(ur)
	userHandlerNonDependency := controller.NewUserController(us)

	// Wire go

	userRouterPublic := Router.Group("/user")
	{
		userRouterPublic.POST("/register", userHandlerNonDependency.Register) // register --> yes --> no
		userRouterPublic.POST("/otp")
	}

	//private router
	userRouterPrivate := Router.Group("/admin/user")
	//userRouterPrivate.Use(limiter())
	//userRouterPrivate.Use(Authen())
	//userRouterPrivate.Use(Permission())
	{
		userRouterPrivate.POST("/active_user")
	}
}
