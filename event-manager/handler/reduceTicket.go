package handler

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/julioceno/ticket-easy/event-manager/config/logger"
	"github.com/julioceno/ticket-easy/event-manager/schemas"
	"github.com/julioceno/ticket-easy/event-manager/utils"
)

func ReduceTicket(ctx *gin.Context) {
	eventId, err := utils.GetIdParam(ctx)
	if err != nil {
		logger.Error("Id not exists", err)
		utils.SendError(ctx, http.StatusBadRequest, "Id não foi especificado")
		return
	}

	ctxMongo, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	event := eventsRepository.FindById(eventId, ctxMongo)
	if hasError := throwErrorIfNotExistsEvent(ctx, event); hasError {
		return
	}

	if hasError := throwErrorIfNotExistsMoreTickets(ctx, event); hasError {
		return
	}

	utils.SendSuccess(ctx, "POST", event)
}

func throwErrorIfNotExistsEvent(ctx *gin.Context, event *schemas.Event) bool {
	if event != nil {
		return false
	}

	eventId, _ := utils.GetIdParam(ctx)
	logger.Error("Event not found", nil)
	utils.SendError(ctx, http.StatusNotFound, fmt.Sprintf("Evento com o id %s não existe", eventId))

	return true
}

func throwErrorIfNotExistsMoreTickets(ctx *gin.Context, event *schemas.Event) bool {
	if existsTickets := event.QuantityTickets > 1; existsTickets {
		return false
	}

	logger.Error("Not exists more tickets", nil)
	utils.SendError(ctx, http.StatusBadRequest, "Evento não tem mais ingressos diponivéis")
	return true
}
