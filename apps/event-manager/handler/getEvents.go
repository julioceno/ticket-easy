package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/julioceno/ticket-easy/apps/event-manager/config/logger"
	"github.com/julioceno/ticket-easy/apps/event-manager/utils"
)

func GetEvents(ctx *gin.Context) {
	ctxMongo, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	count, err := eventsRepository.CountEvents(ctx, ctxMongo)
	if err != nil {
		logger.Error("Error in get cursor event", err)
		utils.SendError(ctx, http.StatusBadRequest, "Não foi possível recuperar os eventos")
		return
	}

	events, err := eventsRepository.FetchEvents(ctx, ctxMongo)
	if err != nil {
		utils.SendError(ctx, http.StatusBadRequest, "Não foi possível recuperar os eventos")
		return
	}

	utils.SendSuccess(utils.SendSuccesStruct{ctx, "GET", utils.ResponseFormat{Count: count, Data: events}, nil})
}
