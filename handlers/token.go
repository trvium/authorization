package handlers

import (
	"github.com/gin-gonic/gin"
)

func TokenHandler(c *gin.Context) {
	token := c.GetHeader("Authorization")
	token = token[len("Bearer "):]
	c.JSON(200, gin.H{"token": token})
}
