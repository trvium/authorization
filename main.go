package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/trvium/authorization/routes"
)

func main() {
	godotenv.Load()
	router := gin.Default()
	routes.SetupRoutes(router)
	router.Run(":8000")
}
