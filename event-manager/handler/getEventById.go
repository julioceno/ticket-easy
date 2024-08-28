package handler

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/julioceno/ticket-easy/event-manager/config/logger"
	"github.com/julioceno/ticket-easy/event-manager/utils"
)

func GetEventById(ctx *gin.Context) {
	id := ctx.Param("id")
	if strings.TrimSpace(id) == "" {
		logger.Error("Occured error when get id", nil)
		utils.SendError(ctx, http.StatusBadRequest, "Ocorreu um erro ao obter o id enviado")
		return
	}

	ctxMongo, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	event := eventsRepository.FindById(id, ctxMongo)
	if &event == nil {
		logger.Error("Event not found", nil)
		utils.SendError(ctx, http.StatusNotFound, fmt.Sprintf("Event with id %s not exists", id))
		return
	}

	utils.SendSuccess(ctx, "GET", event)
}
