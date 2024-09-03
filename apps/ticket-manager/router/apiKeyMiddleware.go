package router

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/julioceno/ticket-easy/apps/ticket-manager/config/logger"
	"github.com/julioceno/ticket-easy/apps/ticket-manager/utils"
)

func apiKeyMiddleware() gin.HandlerFunc {
	apikey := os.Getenv("API_KEY")
	if apikey == "" {
		logger.Fatal("Api key is not valid", nil)
	}

	return func(context *gin.Context) {
		apiKeyReceveid := context.GetHeader("x-api-key")

		if apikey != apiKeyReceveid {
			logger.Info("Api key is invalid")
			utils.SendError(context, http.StatusUnauthorized, "Não Autorizado")
			context.Abort()
			return
		}

		logger.Info("Api key is valid, call next route")
		context.Next()
	}
}
