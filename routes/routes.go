package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/trvium/authorization/handlers"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/info", handlers.GetInfo)
	router.GET("/plan", handlers.FindPlans)
	router.POST("/plan", handlers.CreatePlans)
}
