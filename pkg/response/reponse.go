package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

var (
	ErrInvalidAuthToken = &Error{
		Code:    http.StatusUnauthorized,
		Message: "Invalid token",
	}

	ErrCodeInvalidParams = &Error{
		Code:    http.StatusBadRequest,
		Message: "Invalid parameters",
	}
)

func ResponseError(c *gin.Context, code int, message string) {
	if message == "" {
		message = msg[code]
	}

	c.JSON(http.StatusOK, gin.H{
		"success": false,
		"code":    code,
		"message": message,
	})
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    data,
	})
}
