package router

import (
	"github.com/gin-gonic/gin"
	"github.com/julioceno/ticket-easy/apps/ticket-manager/handler"
)

func initializeRoutes(routes *gin.Engine) {
	ticketRoutes := routes.Group("tickets")

	ticketRoutes.POST("/", handler.CreateTicket)
	ticketRoutes.GET("/:id", handler.GetTicketById)
	ticketRoutes.PATCH("/:id", handler.ReceveidTicketToUpdateStatus)
}
