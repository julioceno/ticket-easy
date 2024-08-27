package router

import "github.com/gin-gonic/gin"

func initializeRoutes(routes *gin.Engine) {
	intializeEventsRoutes(routes)
}
