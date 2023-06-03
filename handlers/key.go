package handlers

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/trvium/authorization/db"
	"github.com/trvium/authorization/models"
	"github.com/trvium/authorization/utils"
)

func ValidateKey(c *gin.Context) {
	input_key := c.GetHeader("x-api-key")

	api_key := &models.ApiKey{}
	err := db.DB.Where("key = ?", input_key).First(api_key).Error
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
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
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	plan := &models.Plan{}
	err = db.DB.Where("id = ?", user.PlanID).First(plan).Error
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
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

func GenerateKey(c *gin.Context) {
	token := c.GetHeader("Authorization")
	token = token[len("Bearer "):]

	claims, err := utils.DecodeToken(token)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	user := &models.User{}
	err = db.DB.Where("email = ?", claims.Email).First(user).Error
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	api_keys := []models.ApiKey{}
	db.DB.Where("user_id = ?", user.ID).Find(&api_keys)

	api_key_max_used := 0
	if len(api_keys) >= 0 {
		for _, api_key := range api_keys {
			if api_key.QuotaUsed > api_key_max_used {
				api_key_max_used = api_key.QuotaUsed
			}
			api_key.Valid = false
			api_key.QuotaUsed = 0
			db.DB.Save(&api_key)
		}
	}

	api_key := &models.ApiKey{
		ID:          utils.GenerateUUID(),
		Key:         utils.GenerateSHA256(),
		UserID:      user.ID,
		Valid:       true,
		QuotaUsed:   api_key_max_used,
		RenewalDate: time.Now().AddDate(0, 1, 0),
	}

	db.DB.Create(&api_key)

	c.JSON(200, gin.H{"data": api_key})
}
