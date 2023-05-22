package handlers

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/gin-gonic/gin"
	"github.com/trvium/authorization/utils"
)

func InfoHandler(c *gin.Context) {
	token := c.GetHeader("Authorization")
	token = token[len("Bearer "):]

	claims, err := utils.DecodeToken(token)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	email := claims.Email
	plan_type := "Enterprise"
	plan_limit := 500
	plan_used := 125

	hash := md5.Sum([]byte(email))
	hash_string := hex.EncodeToString(hash[:])

	c.JSON(200, gin.H{
		"email":      email,
		"plan_type":  plan_type,
		"plan_limit": plan_limit,
		"plan_used":  plan_used,
		"api_key":    hash_string,
	})
}
