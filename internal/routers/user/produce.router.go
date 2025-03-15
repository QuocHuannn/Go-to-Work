package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProduceRouter struct{}

func (pr *ProduceRouter) InitProductRouter(Router *gin.RouterGroup) {
	// public router
	productRouterPublic := Router.Group("/product")
	{
		productRouterPublic.GET("/search", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"status":  "ok",
				"message": "Product search endpoint",
				"data": []map[string]interface{}{
					{"id": 1, "name": "Product 1", "price": 100},
					{"id": 2, "name": "Product 2", "price": 200},
				},
			})
		})

		productRouterPublic.GET("/detail/:id", func(c *gin.Context) {
			id := c.Param("id")
			c.JSON(http.StatusOK, gin.H{
				"status":  "ok",
				"message": "Product detail endpoint",
				"data": map[string]interface{}{
					"id":          id,
					"name":        "Product " + id,
					"price":       100,
					"description": "This is product " + id,
				},
			})
		})
	}
	//
}
