package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/trvium/authorization/handlers"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/hello", handlers.HelloHandler)
	router.GET("/token", handlers.TokenHandler)
	router.GET("/info", handlers.InfoHandler)
}
