package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PongController struct{}

func NewPongController() *PongController {
	return &PongController{}
}

func (P *PongController) Pong(c *gin.Context) {
	name := c.DefaultQuery("name", "Guest")
	uid := c.Query("uid")
	c.JSON(http.StatusOK, gin.H{
		"message": "ping..pong " + name,
		"id":      uid,
		"users":   []string{"user1", "user2"},
	})
}
