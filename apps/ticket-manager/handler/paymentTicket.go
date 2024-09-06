package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/julioceno/ticket-easy/apps/ticket-manager/config/logger"
	"github.com/julioceno/ticket-easy/apps/ticket-manager/schemas"
	"github.com/julioceno/ticket-easy/apps/ticket-manager/utils"
)

func PaymnetTicket(ctx *gin.Context) {
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

	status := schemas.StatusCompleted
	body := _updateStatusTicket{Status: &status}

	ticket, errUpdateTicket := updateTicket(&ctxMongo, &body, ticket)
	if errUpdateTicket != nil {
		utils.SendError(ctx, errUpdateTicket.Code, errUpdateTicket.Message)
		return
	}

	deleteLambdaExpression(ticket.Id.Hex())
	utils.SendSuccess(utils.SendSuccesStruct{ctx, "PATCH", nil, nil})
}
