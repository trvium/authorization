package main

import (
	"github.com/gin-gonic/gin"
	"github.com/trvium/authorization/routes"
)

func main() {
	router := gin.Default()
	routes.SetupRoutes(router)
	router.Run(":8000")
}
