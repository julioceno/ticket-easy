package handler

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/julioceno/ticket-easy/event-manager/config/logger"
	"github.com/julioceno/ticket-easy/event-manager/utils"
)

func GetEventById(ctx *gin.Context) {
	id, err := utils.GetIdParam(ctx)
	if err != nil {
		logger.Error("Id not exists", err)
		utils.SendError(ctx, http.StatusBadRequest, "Id não foi especificado")
		return
	}

	ctxMongo, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	event := eventsRepository.FindById(id, ctxMongo)
	if event == nil {
		logger.Error("Event not found", nil)
		utils.SendError(ctx, http.StatusNotFound, fmt.Sprintf("Evento com o id %s não existe", id))
		return
	}

	utils.SendSuccess(ctx, "GET", event)
}
