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

type SendSuccesStruct struct {
	Ctx    *gin.Context
	Op     string
	Data   interface{}
	Status *int
}

func SendSuccess(params SendSuccesStruct) {
	if params.Status == nil {
		defaultStatus := http.StatusOK
		params.Status = &defaultStatus
	}

	params.Ctx.Header("Content-type", "application/json")
	params.Ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation from handler: %s successfull", params.Op),
		"data":    params.Data,
		"status":  params.Status,
	})
}

func SendError(ctx *gin.Context, code int, msg string) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, gin.H{
		"message": msg,
		"status":  code,
	})
}
