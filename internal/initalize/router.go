package initalize

import (
	"net/http"

	"github.com/QuocHuannn/Go-to-Work/global"
	"github.com/QuocHuannn/Go-to-Work/internal/routers"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	var r *gin.Engine
	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}

	// Load HTML templates
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./static")

	// middleware
	// r.Use() // logging
	// r.Use() // cross
	// r.Use() // limiter global
	userRouter := routers.RouterGroupApp.User

	MainGroup := r.Group("/v1/2024")
	{
		MainGroup.GET("/checkStatus", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"status":  "ok",
				"message": "Service is running",
			})
		})
	}
	{
		userRouter.InitUserRouter(MainGroup)
		userRouter.InitProductRouter(MainGroup)
	}

	// Web routes without v1/2024 prefix
	WebGroup := r.Group("/")
	{
		userRouter.InitUserRouter(WebGroup)
	}

	return r
}
