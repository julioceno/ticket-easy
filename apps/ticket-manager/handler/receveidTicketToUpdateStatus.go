package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/julioceno/ticket-easy/apps/ticket-manager/config/logger"
	"github.com/julioceno/ticket-easy/apps/ticket-manager/utils"
)

func ReceveidTicketToUpdateStatus(ctx *gin.Context) {
	id, body, hasError := getIdAndBody(ctx)
	if hasError {
		return
	}

	if msgError := updateStatusTicket(id, body); msgError != nil {
		utils.SendError(ctx, http.StatusNotFound, *msgError)
		return
	}

	responseStatus := http.StatusNoContent
	utils.SendSuccess(utils.SendSuccesStruct{ctx, "PATCH", nil, &responseStatus})
}

func getIdAndBody(ctx *gin.Context) (*string, *_updateStatusTicket, bool) {
	id, err := utils.GetIdParam(ctx)
	if err != nil {
		logger.Error("Ocurred error when get id", err)
		utils.SendError(ctx, http.StatusNotFound, "Não foi possível obter o id")
		return nil, nil, true
	}

	notHasBodyToDecode := ctx.Request.Body == nil || ctx.Request.ContentLength == 0
	if notHasBodyToDecode {
		return &id, nil, false
	}

	var body _updateStatusTicket
	if err := utils.DecodeBody(ctx, &body); err != nil {
		logger.Error("Ocurred error when try decode body", err)
		utils.SendError(ctx, http.StatusBadRequest, "Ocorreu um erro ao tentar criar o ticket")
		return nil, nil, true
	}

	return &id, &body, false
}
