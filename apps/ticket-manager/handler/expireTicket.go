package handler

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/julioceno/ticket-easy/apps/ticket-manager/config/logger"
	"github.com/julioceno/ticket-easy/apps/ticket-manager/handler/awsServices"
	"github.com/julioceno/ticket-easy/apps/ticket-manager/schemas"
	"github.com/julioceno/ticket-easy/apps/ticket-manager/utils"
)

func ExpireTicket(ctx *gin.Context) {
	id, err := utils.GetIdParam(ctx)
	if err != nil {
		logger.Error("Ocurred error when get id", err)
		utils.SendError(ctx, http.StatusNotFound, "Não foi possível obter o id")
		return
	}

	ctxMongo, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ticket, messageError := getTicket(&id, &ctxMongo, nil)
	if messageError != nil {
		utils.SendError(ctx, http.StatusNotFound, *messageError)
		return
	}

	isExpired := verifyIsExpired(ticket)
	if !isExpired {
		utils.SendSuccess(utils.SendSuccesStruct{ctx, "PATCH", nil, nil})
		return
	}

	needRollbackTicket := ticket.Status == schemas.StatusBuying
	if needRollbackTicket {
		status := schemas.StatusError
		msg := "Ingresso não foi pago no tempo"
		body := _updateStatusTicket{Status: &status, MessageError: &msg}

		if _, err := updateTicket(&ctxMongo, &body, ticket); err != nil {
			utils.SendError(ctx, err.Code, err.Message)
			return
		}

		sendRollbackTicketHttp(&ticket.EventId)
	}

	if err := deleteLambdaExpression(ticket.Id.Hex()); err != nil {
		utils.SendError(ctx, err.Code, err.Message)
		return
	}

	utils.SendSuccess(utils.SendSuccesStruct{ctx, "PATCH", nil, nil})
}

func verifyIsExpired(ticket *schemas.Ticket) bool {
	createdAt := ticket.CreatedAt
	currentTime := time.Now()
	diff := currentTime.Sub(createdAt)
	isExpired := diff > 2*time.Minute

	return isExpired
}

func deleteLambdaExpression(ticketId string) *utils.ErrorPattern {
	if err := awsServices.DeleteEvent(ticketId); err != nil {
		errorCreated := utils.ErrorPattern{
			Code:    http.StatusBadRequest,
			Message: "Ocorreu um erro ao tentar garantir o ingresso",
		}
		return &errorCreated
	}
	return nil
}

func sendRollbackTicketHttp(eventId *string) bool {
	eventUrl := os.Getenv("EVENT_URL")
	apiKey := os.Getenv("EVENT_API_KEY")
	url := fmt.Sprintf("%s/events/%s/rollback-ticket", eventUrl, *eventId)

	req, err := http.NewRequest("PATCH", url, nil)
	if err != nil {
		logger.Error("Occurred error in build request", err)
		return true
	}

	req.Header.Set("x-api-key", apiKey)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		logger.Error("Occurred error in call event system", err)
		return true
	}

	if ocurredAnyError := response.StatusCode != http.StatusNoContent; ocurredAnyError {
		return true
	}

	return false
}
