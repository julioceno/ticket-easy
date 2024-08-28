package utils

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseFormat struct {
	Count int64       `json:"count"`
	Data  interface{} `json:"data"`
}

func SendSuccess(ctx *gin.Context, op string, data interface{}) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation from handler: %s successfull", op),
		"data":    data,
		"status":  200,
	})
}

func SendError(ctx *gin.Context, code int, msg string) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, gin.H{
		"message": msg,
		"status":  code,
	})
}
