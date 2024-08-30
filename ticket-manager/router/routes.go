package router

import (
	"github.com/gin-gonic/gin"
	"github.com/julioceno/ticket-easy/ticket-manager/handler"
)

func initializeRoutes(routes *gin.Engine) {
	ticketRoutes := routes.Group("tickets")

	ticketRoutes.POST("/", handler.CreateTicket)
}
