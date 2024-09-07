package handler

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/julioceno/ticket-easy/apps/ticket-manager/config/logger"
	"github.com/julioceno/ticket-easy/apps/ticket-manager/schemas"
	"github.com/julioceno/ticket-easy/apps/ticket-manager/utils"
	"go.mongodb.org/mongo-driver/bson"
)

func GetTicketById(ctx *gin.Context) {
	id, userId, hasError := getIdAndUserId(ctx)
	if hasError {
		return
	}
	ctxMongo, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filters := bson.M{"userId": userId}
	ticket := ticketsRepository.FindById(id, &ctxMongo, filters)
	if hasError := throwErrorIfNotExistsEvent(ctx, ticket); hasError {
		return
	}

	response := ticket.ToResponse()
	utils.SendSuccess(utils.SendSuccesStruct{ctx, "GET", response, nil})
}

func throwErrorIfNotExistsEvent(ctx *gin.Context, ticket *schemas.Ticket) bool {
	if ticket != nil {
		return false
	}

	ticketId, _ := utils.GetIdParam(ctx)
	logger.Error("Event not found", nil)
	utils.SendError(ctx, http.StatusNotFound, fmt.Sprintf("Ingresso com o id %s não existe", ticketId))

	return true
}
