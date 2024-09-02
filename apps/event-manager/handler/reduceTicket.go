package handler

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/julioceno/ticket-easy/apps/event-manager/config/logger"
	"github.com/julioceno/ticket-easy/apps/event-manager/schemas"
	"github.com/julioceno/ticket-easy/apps/event-manager/utils"
)

type reduceTicketBody struct {
	TicketId string `json:"ticketId" validate:"required"`
}

func ReduceTicket(ctx *gin.Context) {
	eventId, body, hasError := getIdAndBody(ctx)
	if hasError {
		return
	}

	ctxMongo, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	event := eventsRepository.FindById(eventId, &ctxMongo)
	if hasError := throwErrorIfNotExistsEvent(ctx, event); hasError {
		return
	}

	if hasError := throwErrorIfNotExistsMoreTickets(ctx, event); hasError {
		return
	}

	utils.SendSuccess(ctx, "POST", event)
	go decreaseTicket(event, &body.TicketId)
}

func getIdAndBody(ctx *gin.Context) (*string, *reduceTicketBody, bool) {
	eventId, err := utils.GetIdParam(ctx)
	if err != nil {
		logger.Error("Id not exists", err)
		utils.SendError(ctx, http.StatusBadRequest, "Id não foi especificado")
		return nil, nil, true
	}

	var body reduceTicketBody
	if err := utils.DecodeBody(ctx, &body); err != nil {
		logger.Error("Ocurred error when try decode body", err)
		utils.SendError(ctx, http.StatusBadRequest, "Ocorreu um erro ao tentar garantir o ingresso")
		return nil, nil, true
	}

	return &eventId, &body, false
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
	if existsTickets := event.QuantityTickets > 0; existsTickets {
		return false
	}

	logger.Error("Not exists more tickets", nil)
	utils.SendError(ctx, http.StatusBadRequest, "Evento não tem mais ingressos diponivéis")
	return true
}
