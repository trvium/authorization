package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/trvium/authorization/db"
	"github.com/trvium/authorization/models"
	"github.com/trvium/authorization/utils"
)

func FindPlans(c *gin.Context) {
	var plans []models.Plan
	db.DB.Find(&plans)

	c.JSON(http.StatusOK, gin.H{"data": plans})
}

func CreatePlans(c *gin.Context) {
	var plans []models.Plan
	db.DB.Find(&plans)

	if len(plans) > 0 {
		c.JSON(http.StatusOK, gin.H{"data": plans})
		return
	}

	db.DB.Create(&models.Plan{Name: "Hobby", Quota: 50, ID: utils.GenerateUUID()})
	db.DB.Create(&models.Plan{Name: "Business", Quota: 250, ID: utils.GenerateUUID()})
	db.DB.Create(&models.Plan{Name: "Enterprise", Quota: 500, ID: utils.GenerateUUID()})

	db.DB.Find(&plans)

	c.JSON(http.StatusOK, gin.H{"data": plans})
}
