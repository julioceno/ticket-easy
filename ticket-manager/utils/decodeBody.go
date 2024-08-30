package utils

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func DecodeBody(ctx *gin.Context, body interface{}) error {
	if err := json.NewDecoder(ctx.Request.Body).Decode(&body); err != nil {
		return err
	}

	return nil
}
