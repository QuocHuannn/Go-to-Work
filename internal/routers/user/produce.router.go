package user

import "github.com/gin-gonic/gin"

type ProduceRouter struct{}

func (pr *ProduceRouter) InitProductRouter(Router *gin.RouterGroup) {
	// public router
	productRouterPublic := Router.Group("/product")
	{
		productRouterPublic.GET("/search")
		productRouterPublic.GET("/detail/:id")
	}
	//
}
