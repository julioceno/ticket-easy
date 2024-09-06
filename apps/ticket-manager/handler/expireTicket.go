package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/julioceno/ticket-easy/apps/ticket-manager/config/logger"
	"github.com/julioceno/ticket-easy/apps/ticket-manager/handler/awsServices"
	"github.com/julioceno/ticket-easy/apps/ticket-manager/schemas"
	"github.com/julioceno/ticket-easy/apps/ticket-manager/utils"
)

func ExpireTicket(ctx *gin.Context) {
	id, err := utils.GetIdParam(ctx)
	if err != nil {
		logger.Error("Ocurred error when get id", err)
		utils.SendError(ctx, http.StatusNotFound, "Não foi possível obter o id")
		return
	}

	ctxMongo, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ticket, messageError := getTicket(&id, &ctxMongo)
	if messageError != nil {
		utils.SendError(ctx, http.StatusNotFound, *messageError)
		return
	}

	isExpired := verifyIsExpired(ticket)
	if !isExpired {
		utils.SendSuccess(utils.SendSuccesStruct{ctx, "PATCH", nil, nil})
		return
	}

	status := schemas.StatusError
	msg := "Ingresso não foi pago no tempo"
	body := _updateStatusTicket{Status: &status, MessageError: &msg}
	if _, errUpdateTicket := updateTicket(&ctxMongo, &body, ticket); errUpdateTicket != nil {
		utils.SendError(ctx, errUpdateTicket.Code, errUpdateTicket.Message)
		return
	}

	deleteLambdaExpression(ticket.Id.Hex())
	utils.SendSuccess(utils.SendSuccesStruct{ctx, "PATCH", nil, nil})
}

func verifyIsExpired(ticket *schemas.Ticket) bool {
	createdAt := ticket.CreatedAt
	currentTime := time.Now()
	diff := currentTime.Sub(createdAt)
	isExpired := diff > 2*time.Minute

	return isExpired
}

func deleteLambdaExpression(ticketId string) *utils.ErrorPattern {
	if err := awsServices.DeleteEvent(ticketId); err != nil {
		errorCreated := utils.ErrorPattern{
			Code:    http.StatusBadRequest,
			Message: "Ocorreu um erro ao tentar garantir o ingresso",
		}
		return &errorCreated
	}
	return nil
}
