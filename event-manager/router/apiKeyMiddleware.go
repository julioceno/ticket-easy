package router

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/julioceno/ticket-easy/event-manager/config/logger"
	"github.com/julioceno/ticket-easy/event-manager/utils"
)

func apiKeyMiddleware() gin.HandlerFunc {
	err := godotenv.Load()
	if err != nil {
		logger.Fatal("Error loading .env file", err)
	}

	apikey := os.Getenv("API_KEY")

	return func(context *gin.Context) {
		apiKeyReceveid := context.GetHeader("x-api-key")

		if apikey != apiKeyReceveid {
			logger.Info("Api key is invalid")
			utils.SendError(context, http.StatusUnauthorized, "Não Autorizado")
			return
		}

		logger.Info("Api key is valid, call next route")
		context.Next()
	}
}
