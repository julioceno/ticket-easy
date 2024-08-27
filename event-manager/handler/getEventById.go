package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/julioceno/ticket-easy/event-manager/utils"
)

func GetEventById(ctx *gin.Context) {
	utils.SendSuccess(ctx, "GET", "")
}
