package handler

import (
	"context"
	"errors"
	"fmt"

	"github.com/julioceno/ticket-easy/apps/ticket-manager/config/logger"
	"github.com/julioceno/ticket-easy/apps/ticket-manager/schemas"
)

func updateEventError(ctxMongo *context.Context, messageError *string, ticket *schemas.Ticket) (*schemas.Ticket, error) {
	ticket.MessageError = *messageError
	ticket.Status = schemas.StatusError

	id := ticket.Id.Hex()
	ticket, err := ticketsRepository.Update(&id, ctxMongo, ticket)

	return ticket, err
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
