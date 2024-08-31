package router

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/julioceno/ticket-easy/event-manager/config/logger"
)

func Initialize() {
	if err := godotenv.Load(); err != nil {
		logger.Fatal("Error loading .env file", err)
	}

	router := gin.Default()
	router.Use(apiKeyMiddleware())

	initializeRoutes(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = ":8002"
	}

	router.Run(port)
}
