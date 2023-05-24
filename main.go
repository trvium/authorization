package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/trvium/authorization/routes"
)

func main() {
	godotenv.Load()

	router := gin.Default()
	routes.SetupRoutes(router)

	port := ":" + os.Getenv("PORT")
	if port == ":" {
		port = ":8000"
	}
	router.Run(port)
}
