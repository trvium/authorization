package handlers

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/trvium/authorization/db"
	"github.com/trvium/authorization/models"
	"github.com/trvium/authorization/utils"
	"gorm.io/gorm"
)

func GetInfo(c *gin.Context) {
	token := c.GetHeader("Authorization")
	token = token[len("Bearer "):]

	claims, err := utils.DecodeToken(token)
	if err != nil {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	email := claims.Email

	user := &models.User{}
	err = db.DB.Where("email = ?", email).First(user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {

			planName := "Hobby"

			plan := &models.Plan{}
			err = db.DB.Where("name = ?", planName).First(plan).Error
			if err != nil {
				return
			}

			userEntity := &models.User{
				ID:     utils.GenerateUUID(),
				PlanID: plan.ID,
				Email:  email,
			}

			err = db.DB.Create(userEntity).Error
			if err != nil {
				return
			}

			user = userEntity
		} else {
			return
		}
	}

	plan := &models.Plan{}
	err = db.DB.Where("id = ?", user.PlanID).First(plan).Error
	if err != nil {
		return
	}

	apiKey := &models.ApiKey{}
	err = db.DB.Where("user_id = ? AND valid = ?", user.ID, true).First(apiKey).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			apiKeyEntity := &models.ApiKey{
				ID:          utils.GenerateUUID(),
				UserID:      user.ID,
				Key:         utils.GenerateSHA256(),
				Valid:       true,
				QuotaUsed:   0,
				RenewalDate: time.Now().AddDate(0, 1, 0),
			}

			err = db.DB.Create(apiKeyEntity).Error
			if err != nil {
				return
			}

			apiKey = apiKeyEntity
		} else {
			return
		}
	}

	plan_type := plan.Name
	plan_limit := plan.Quota
	plan_used := apiKey.QuotaUsed
	api_key := apiKey.Key

	c.JSON(200, gin.H{
		"email":      email,
		"plan_type":  plan_type,
		"plan_limit": plan_limit,
		"plan_used":  plan_used,
		"api_key":    api_key,
	})
}
