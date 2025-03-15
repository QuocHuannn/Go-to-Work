package middleware

import (
	"github.com/QuocHuannn/Go-to-Work/pkg/response"
	"github.com/gin-gonic/gin"
)

func AuthenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token != "valid-token" {
			response.ResponseError(c, response.ErrInvalidAuthToken.Code, response.ErrInvalidAuthToken.Message)
			c.Abort()
			return
		}
		c.Next()
	}
}
