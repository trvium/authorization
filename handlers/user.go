package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/trvium/authorization/db"
	"github.com/trvium/authorization/models"
	"github.com/trvium/authorization/utils"
)

func ChangePlan(c *gin.Context) {
	token := c.GetHeader("Authorization")
	token = token[len("Bearer "):]

	claims, err := utils.DecodeToken(token)
	if err != nil {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	user := &models.User{}
	err = db.DB.Where("email = ?", claims.Email).First(user).Error
	if err != nil {
		c.JSON(401, gin.H{"error": "User not found"})
		return
	}

	plan := &models.Plan{}
	err = db.DB.Where("id = ?", c.Param("id")).First(plan).Error
	if err != nil {
		c.JSON(401, gin.H{"error": "Plan not found"})
		return
	}

	user.PlanID = plan.ID
	db.DB.Save(user)

	c.JSON(200, gin.H{"data": user})
}
