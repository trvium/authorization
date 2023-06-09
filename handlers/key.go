package handlers

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/trvium/authorization/db"
	"github.com/trvium/authorization/models"
)

func ValidateKey(c *gin.Context) {
	input_key := c.GetHeader("x-api-key")

	api_key := &models.ApiKey{}
	err := db.DB.Where("key = ?", input_key).First(api_key).Error
	if err != nil {
		c.JSON(401, gin.H{"error": "Invalid API Key"})
		return
	}

	if !api_key.Valid {
		c.JSON(401, gin.H{"error": "Invalid API Key"})
		return
	}

	if api_key.RenewalDate.Before(time.Now()) || api_key.RenewalDate.Equal(time.Now()) {
		api_key.RenewalDate = time.Now().AddDate(0, 1, 0)
		api_key.QuotaUsed = 0
		db.DB.Save(api_key)
	}

	user := &models.User{}
	err = db.DB.Where("id = ?", api_key.UserID).First(user).Error
	if err != nil {
		c.JSON(401, gin.H{"error": "Invalid API Key"})
		return
	}

	plan := &models.Plan{}
	err = db.DB.Where("id = ?", user.PlanID).First(plan).Error
	if err != nil {
		c.JSON(401, gin.H{"error": "Invalid API Key"})
		return
	}

	if api_key.QuotaUsed >= plan.Quota {
		c.JSON(401, gin.H{"error": "Quota Exceeded"})
		return
	}

	api_key.QuotaUsed += 1
	db.DB.Save(api_key)

	c.JSON(200, gin.H{"data": "Valid API Key"})
}
