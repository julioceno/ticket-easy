package handler

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/julioceno/ticket-easy/apps/ticket-manager/config/logger"
	"github.com/julioceno/ticket-easy/apps/ticket-manager/schemas"
	"github.com/julioceno/ticket-easy/apps/ticket-manager/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func updateEventError(ctxMongo *context.Context, messageError *string, ticket *schemas.Ticket) (*schemas.Ticket, error) {
	ticket.MessageError = *messageError
	ticket.Status = schemas.StatusError

	id := ticket.Id.Hex()
	ticket, err := ticketsRepository.Update(&id, ctxMongo, ticket)

	return ticket, err
}

func getTicket(id *string, ctxMongo *context.Context, filters primitive.M) (*schemas.Ticket, *string) {
	ticket := ticketsRepository.FindById(id, ctxMongo, filters)
	if ticket == nil {
		logger.Error("When try update ticket, document not exists", errors.New(fmt.Sprint("Ticket with id %v not exists", id)))
		msg := fmt.Sprintf("Ticket de id %v não existe", *id)
		return nil, &msg
	}

	return ticket, nil
}

func getIdAndUserId(ctx *gin.Context) (*string, *string, bool) {
	id, err := utils.GetIdParam(ctx)
	if err != nil {
		logger.Error("Id not exists", err)
		utils.SendError(ctx, http.StatusBadRequest, "Id não foi especificado")
		return nil, nil, true
	}

	userId, isOk := ctx.GetQuery("userId")
	if !isOk {
		errCreated := errors.New("User Id not exists")
		logger.Error("User Id not exists", errCreated)
		utils.SendError(ctx, http.StatusBadRequest, "UserId não foi especificado")
		return nil, nil, true
	}

	return &id, &userId, false
}
