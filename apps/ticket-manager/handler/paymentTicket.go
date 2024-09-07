package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/julioceno/ticket-easy/apps/ticket-manager/schemas"
	"github.com/julioceno/ticket-easy/apps/ticket-manager/utils"
	"go.mongodb.org/mongo-driver/bson"
)

func PaymnetTicket(ctx *gin.Context) {
	id, userId, hasError := getIdAndUserId(ctx)
	if hasError {
		return
	}

	ctxMongo, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filters := bson.M{"userId": userId}
	ticket, messageError := getTicket(id, &ctxMongo, filters)
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

	go deleteEventBridge(ticket.Id.Hex())
	utils.SendSuccess(utils.SendSuccesStruct{ctx, "PATCH", nil, nil})
}
