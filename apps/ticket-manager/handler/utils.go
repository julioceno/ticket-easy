package handler

import (
	"context"

	"github.com/julioceno/ticket-easy/apps/ticket-manager/schemas"
)

func updateEventError(ctxMongo *context.Context, messageError *string, ticket *schemas.Ticket) {
	// TODO: salvar mensagens de erro mais concretas. Por exemplo, quando os ingressos de esgotarem, retornar uma mensagem dizendo que os ingressos esgotaram
	ticket.MessageError = *messageError
	ticket.Status = schemas.StatusError

	id := ticket.Id.Hex()
	ticketsRepository.Update(&id, ctxMongo, ticket)
}
