package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/julioceno/ticket-easy/ticket-manager/config/logger"
	"github.com/julioceno/ticket-easy/ticket-manager/utils"
)

func GetTicketById(ctx *gin.Context) {
	id, err := utils.GetIdParam(ctx)
	if err != nil {
		logger.Error("Id not exists", err)
		utils.SendError(ctx, http.StatusBadRequest, "Id não foi especificado")
		return
	}

	ctxMongo, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ticket := ticketsRepository.FindById(id, ctxMongo)
	if ticket == nil {
		logger.Error("Ticket not exists", err)
		utils.SendError(ctx, http.StatusNotFound, "Não existe um ticket com esse id")
		return
	}

	response := ticket.ToResponse()
	utils.SendSuccess(utils.SendSuccesStruct{ctx, "POST", response, nil})
}
