package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/trvium/authorization/handlers"
)

func SetupRoutes(router *gin.Engine) {
	apiGroup := router.Group("/")

	// Health check
	apiGroup.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})

	apiGroup.GET("/plan", handlers.FindPlans)
	apiGroup.POST("/plan", handlers.CreatePlans)
	apiGroup.GET("/info", handlers.GetInfo)
	apiGroup.POST("/key/validate", handlers.ValidateKey)
	apiGroup.PUT("/user/plan/:id", handlers.ChangePlan)
}
