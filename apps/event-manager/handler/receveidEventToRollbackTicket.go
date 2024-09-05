package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/julioceno/ticket-easy/apps/event-manager/config/logger"
	"github.com/julioceno/ticket-easy/apps/event-manager/utils"
)

func ReceveidEventToRollbackTicket(ctx *gin.Context) {
	id, err := utils.GetIdParam(ctx)
	if err != nil {
		logger.Error("Id not exists", err)
		utils.SendError(ctx, http.StatusBadRequest, "Id não foi especificado")
		return
	}

	if _, err := rollbackTicket(id); err != nil {
		utils.SendError(ctx, err.Code, err.Message)
		return
	}

	statusCode := http.StatusNoContent
	utils.SendSuccess(utils.SendSuccesStruct{ctx, "PATCH", nil, &statusCode})
}
