package router

import (
	"github.com/gin-gonic/gin"
	"github.com/julioceno/ticket-easy/event-manager/handler"
)

func intializeEventsRoutes(router *gin.Engine) {
	eventsRoutes := router.Group("/events")

	eventsRoutes.GET("/", handler.GetEvents)
}
