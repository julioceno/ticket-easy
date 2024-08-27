package router

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/julioceno/ticket-easy/event-manager/handler"
)

func Initialize() {
	router := gin.Default()

	handler.IntializeHandler()
	initializeRoutes(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = ":8081"
	}

	router.Run(port)
}
