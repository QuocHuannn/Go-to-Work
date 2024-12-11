package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

var ErrInvalidAuthToken = &Error{
	Code:    http.StatusUnauthorized,
	Message: "Invalid token",
}

func ResponseError(c *gin.Context, err *Error, msg string) {
	c.JSON(err.Code, gin.H{
		"code":    err.Code,
		"message": msg,
	})
}
