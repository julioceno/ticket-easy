package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/julioceno/ticket-easy/apps/ticket-manager/config/logger"
	"github.com/julioceno/ticket-easy/apps/ticket-manager/handler/awsServices"
	"github.com/julioceno/ticket-easy/apps/ticket-manager/schemas"
	"github.com/julioceno/ticket-easy/apps/ticket-manager/utils"
)

type _updateStatusTicket struct {
	MessageError *string               `json:"messageError"`
	Status       *schemas.TicketStatus `json:"status"`
}

func updateStatusTicket(id *string, body *_updateStatusTicket) *utils.ErrorPattern {
	ctxMongo, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ticket, messageError := getTicket(id, &ctxMongo)
	if messageError != nil {
		errorCreated := utils.ErrorPattern{
			Code:    http.StatusNotFound,
			Message: *messageError,
		}

		return &errorCreated
	}

	ticketUpdated, err := updateTicket(&ctxMongo, body, ticket)
	if err != nil {
		return err
	}

	return createLambdaExpression(ticketUpdated.Id.Hex())
}

func updateTicket(ctxMongo *context.Context, body *_updateStatusTicket, ticket *schemas.Ticket) (*schemas.Ticket, *utils.ErrorPattern) {
	defaultError := utils.ErrorPattern{
		Code:    http.StatusNotFound,
		Message: "Ocorreu um erro ao tentar atualizar o ingresso",
	}

	if body != nil && body.MessageError != nil {
		ticketUpdated, err := updateEventError(ctxMongo, body.MessageError, ticket)
		if err != nil {
			return nil, &defaultError
		}
		return ticketUpdated, nil
	}

	id := ticket.Id.Hex()
	ticket.Status = *body.Status
	ticketUpdated, err := ticketsRepository.Update(&id, ctxMongo, ticket)

	if err != nil {
		logger.Error("Ocurred error when try update ticket", err)
		return nil, &defaultError
	}

	return ticketUpdated, nil
}

func createLambdaExpression(ticketId string) *utils.ErrorPattern {
	if err := awsServices.CreateEvent(ticketId); err != nil {
		errorCreated := utils.ErrorPattern{
			Code:    http.StatusBadRequest,
			Message: "Ocorreu um erro ao tentar garantir o ingresso",
		}
		return &errorCreated
	}

	return nil
}
