package handler

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/julioceno/ticket-easy/apps/ticket-manager/config/logger"
	"github.com/julioceno/ticket-easy/apps/ticket-manager/handler/awsServices"
	"github.com/julioceno/ticket-easy/apps/ticket-manager/schemas"
)

type ticketAction struct {
	TicketId     *string `json:"ticketId"`
	MessageError *string `json:"messageError"`
}

func startConsumerDecreaseTicket() {
	for {
		msgResult, err := awsServices.ReceveidMessage(awsServices.QueueDecreaseTicket)
		if err != nil {
			logger.Error("Ocurred error when trying to get queue messages", err)
			time.Sleep(5 * time.Second)
			continue
		}

		consumeReceveidDecreaseTicket(msgResult)
	}
}

func consumeReceveidDecreaseTicket(msgResult *sqs.ReceiveMessageOutput) {
	for _, msg := range msgResult.Messages {
		var result ticketAction

		if err := json.Unmarshal([]byte(*msg.Body), &result); err != nil {
			fmt.Println("Message:", *msg.Body)
			if e, ok := err.(*json.SyntaxError); ok {
				fmt.Printf("Syntax error byte offset %d: %v\n", e.Offset, err)
			} else {
				fmt.Printf("Error deserializing message: %v\n", err)
			}
			continue
		}

		status := schemas.StatusBuying
		body := _updateStatusTicket{
			MessageError: result.MessageError,
			Status:       &status,
		}

		if err := updateStatusTicket(result.TicketId, &body); err != nil {
			logger.Error(fmt.Sprintf("Ocurred error when try update status ticket %v", err.Message), nil)
			continue
		}

		if err := awsServices.DeleteMessage(awsServices.QueueDecreaseTicket, msg.ReceiptHandle); err != nil {
			logger.Error("Ocurred errro when try delete message queue", err)
		}
	}
}
