package router

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/julioceno/ticket-easy/apps/event-manager/config/logger"
	"github.com/julioceno/ticket-easy/apps/event-manager/utils"
)

func apiKeyMiddleware() gin.HandlerFunc {
	apikey := os.Getenv("API_KEY")
	if apikey == "" {
		logger.Fatal("Api key is not valid", nil)
	}

	return func(context *gin.Context) {
		apiKeyReceveid := context.GetHeader("x-api-key")

		if apikey != apiKeyReceveid {
			utils.SendError(context, http.StatusUnauthorized, "Não Autorizado")
			context.Abort()
			return
		}

		context.Next()
	}
}
