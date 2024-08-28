package router

import (
	"os"

	"github.com/gin-gonic/gin"
)

func Initialize() {
	router := gin.Default()
	router.Use(apiKeyMiddleware())

	initializeRoutes(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = ":8081"
	}

	router.Run(port)
}
