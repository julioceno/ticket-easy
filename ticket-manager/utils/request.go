package utils

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetIdParam(ctx *gin.Context) (string, error) {
	id := ctx.Param("id")
	if strings.TrimSpace(id) == "" {
		return "", errors.New("Id param not exist")
	}

	return id, nil
}
