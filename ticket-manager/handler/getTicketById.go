package handler

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/julioceno/ticket-easy/ticket-manager/config/logger"
	"github.com/julioceno/ticket-easy/ticket-manager/schemas"
	"github.com/julioceno/ticket-easy/ticket-manager/utils"
)

// TODO: so pegar o ticket se pertencer ao usuario que for especificado
func GetTicketById(ctx *gin.Context) {
	id, err := utils.GetIdParam(ctx)
	if err != nil {
		logger.Error("Id not exists", err)
		utils.SendError(ctx, http.StatusBadRequest, "Id não foi especificado")
		return
	}

	ctxMongo, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ticket := ticketsRepository.FindById(&id, &ctxMongo)
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
