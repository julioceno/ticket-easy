package router

import (
	"github.com/gin-gonic/gin"
	"github.com/julioceno/ticket-easy/apps/event-manager/handler"
)

func initializeRoutes(routes *gin.Engine) {
	eventsRoutes := routes.Group("/events")

	eventsRoutes.GET("/", handler.GetEvents)
	eventsRoutes.GET("/:id", handler.GetEventById)
	eventsRoutes.POST("/:id/reduce-ticket", handler.ReduceTicket)
}
