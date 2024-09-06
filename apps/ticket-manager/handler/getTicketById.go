package handler

import (
	"context"
	"errors"
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

func throwErrorIfNotExistsEvent(ctx *gin.Context, ticket *schemas.Ticket) bool {
	if ticket != nil {
		return false
	}

	ticketId, _ := utils.GetIdParam(ctx)
	logger.Error("Event not found", nil)
	utils.SendError(ctx, http.StatusNotFound, fmt.Sprintf("Ingresso com o id %s não existe", ticketId))

	return true
}
