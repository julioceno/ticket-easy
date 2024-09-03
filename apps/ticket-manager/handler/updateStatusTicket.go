package handler

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/julioceno/ticket-easy/apps/ticket-manager/config/logger"
	"github.com/julioceno/ticket-easy/apps/ticket-manager/schemas"
)

type _updateStatusTicket struct {
	MessageError *string `json:"messageError"`
}

func updateStatusTicket(id *string, body *_updateStatusTicket) *string {
	ctxMongo, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	ticket, messageError := getTicket(id, &ctxMongo)
	if messageError != nil {
		return messageError
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

	ticket.Status = schemas.StatusBuying
	id := ticket.Id.Hex()
	ticketsRepository.Update(&id, ctxMongo, ticket)
}
