package handler

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/julioceno/ticket-easy/ticket-manager/config/logger"
	"github.com/julioceno/ticket-easy/ticket-manager/schemas"
	"github.com/julioceno/ticket-easy/ticket-manager/utils"
)

type _ReceveidTicketbody struct {
	MessageError *string `json:"messageError"`
}

func ReceveIdResultReduceTicket(ctx *gin.Context) {
	body, hasError := getBody(ctx)
	if hasError {
		return
	}

	ctxMongo, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	ticket, hasError := getTicket(ctx, &ctxMongo)
	if hasError {
		return
	}

	updateTicket(&ctxMongo, body, ticket)
	responseStatus := http.StatusNoContent
	utils.SendSuccess(utils.SendSuccesStruct{ctx, "PATCH", nil, &responseStatus})
}

func getBody(ctx *gin.Context) (*_ReceveidTicketbody, bool) {
	notHasBodyToDecode := ctx.Request.Body == nil || ctx.Request.ContentLength == 0
	if notHasBodyToDecode {
		return nil, false
	}

	var body _ReceveidTicketbody
	if err := utils.DecodeBody(ctx, &body); err != nil {
		logger.Error("Ocurred error when try decode body", err)
		utils.SendError(ctx, http.StatusBadRequest, "Ocorreu um erro ao tentar criar o ticket")
		return nil, true
	}

	return &body, false
}

func getTicket(ctx *gin.Context, ctxMongo *context.Context) (*schemas.Ticket, bool) {
	id, err := utils.GetIdParam(ctx)
	if err != nil {
		logger.Error("When try update ticket, ocurred error when get id", err)
		utils.SendError(ctx, http.StatusNotFound, "Não foi possível obter o id")
	}

	ticket := ticketsRepository.FindById(&id, ctxMongo)
	if ticket == nil {
		logger.Error("When try update ticket, document not exists", errors.New(fmt.Sprint("Ticket with id %s not exists", id)))
		utils.SendError(ctx, http.StatusNotFound, fmt.Sprintf("Ticket de id %s não existe", id))
		return nil, true
	}

	return ticket, false
}

func updateTicket(ctxMongo *context.Context, body *_ReceveidTicketbody, ticket *schemas.Ticket) {
	if body != nil && body.MessageError != nil {
		updateEventError(ctxMongo, body.MessageError, ticket)
		return
	}

	ticket.Status = schemas.StatusBuying
	id := ticket.Id.Hex()
	ticketsRepository.Update(&id, ctxMongo, ticket)
}
