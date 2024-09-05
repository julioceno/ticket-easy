package handler

import (
	"context"
	"errors"
	"fmt"
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
	if err := awsServices.CreateEvent(*id); err != nil {
		errorCreated := utils.ErrorPattern{
			Code:    http.StatusBadRequest,
			Message: "Ocorreu um erro ao tentar garantir o ingresso",
		}
		return &errorCreated
	}

	ctxMongo, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	ticket, messageError := getTicket(id, &ctxMongo)
	if messageError != nil {
		errorCreated := utils.ErrorPattern{
			Code:    http.StatusNotFound,
			Message: *messageError,
		}

		return &errorCreated
	}

	updateTicket(&ctxMongo, body, ticket)
	return nil
}

func getTicket(id *string, ctxMongo *context.Context) (*schemas.Ticket, *string) {
	ticket := ticketsRepository.FindById(id, ctxMongo)
	if ticket == nil {
		logger.Error("When try update ticket, document not exists", errors.New(fmt.Sprint("Ticket with id %v not exists", id)))
		msg := fmt.Sprintf("Ticket de id %v não existe", id)
		return nil, &msg
	}

	return ticket, nil
}

func updateTicket(ctxMongo *context.Context, body *_updateStatusTicket, ticket *schemas.Ticket) {
	if body != nil && body.MessageError != nil {
		updateEventError(ctxMongo, body.MessageError, ticket)
		return
	}

	id := ticket.Id.Hex()
	ticket.Status = *body.Status
	ticketsRepository.Update(&id, ctxMongo, ticket)
}
