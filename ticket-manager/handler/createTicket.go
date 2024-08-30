package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/julioceno/ticket-easy/ticket-manager/config/logger"
	"github.com/julioceno/ticket-easy/ticket-manager/schemas"
	"github.com/julioceno/ticket-easy/ticket-manager/utils"
)

type _body struct {
	UserId  string `json:"userId" validate:"required"`
	EventId string `json:"eventId" validate:"required"`
}

type _response struct {
	UserId  string `json:"userId" validate:"required"`
	EventId string `json:"eventId" validate:"required"`
}

func CreateTicket(ctx *gin.Context) {
	var body _body
	if err := utils.DecodeBody(ctx, &body); err != nil {
		logger.Error("Ocurred error when try decode body", err)
		utils.SendError(ctx, http.StatusBadRequest, "Ocorreu um erro ao tentar criar o ticket")
		return
	}

	ticket, err := createTicket(body)
	if err != nil {
		logger.Error("Ocurred error when format body to create ticket", err)
		utils.SendError(ctx, http.StatusBadRequest, "Ocorreu um erro ao tentar criar o ticket")
		return
	}

	ctxMongo, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ticketCreated, err := ticketsRepository.Create(ctx, ctxMongo, ticket)

	response := ticketCreated.ToResponse()
	responseStatus := http.StatusCreated
	utils.SendSuccess(utils.SendSuccesStruct{ctx, "POST", response, &responseStatus})
}

func createTicket(body _body) (*schemas.Ticket, error) {
	key, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	ticket := schemas.Ticket{
		Status:  schemas.StatusBuy,
		EventId: body.EventId,
		UserId:  body.UserId,
		Key:     key.String(),
	}

	return &ticket, nil
}
