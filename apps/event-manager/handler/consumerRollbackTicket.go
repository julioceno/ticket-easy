package handler

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/julioceno/ticket-easy/apps/event-manager/config/logger"
	"github.com/julioceno/ticket-easy/apps/event-manager/handler/queue"
)

type eventAction struct {
	EventId *string `json:"eventId"`
}

func startConsumerRollbackTicket() {
	for {
		msgResult, err := queue.ReceveidMessage()
		if err != nil {
			logger.Error("Ocurred error when try to get queue messages", err)
			time.Sleep(5 * time.Second)
			continue
		}

		consumeReceveidDecreaseTicket(msgResult)
	}
}

func consumeReceveidDecreaseTicket(msgResult *sqs.ReceiveMessageOutput) {
	for _, msg := range msgResult.Messages {
		var result eventAction

		if err := json.Unmarshal([]byte(*msg.Body), &result); err != nil {
			// TODO: traduzir pra ingles
			fmt.Println("Message:", *msg.Body)
			if e, ok := err.(*json.SyntaxError); ok {
				fmt.Printf("Erro de sintaxe no byte offset %d: %v\n", e.Offset, err)
			} else {
				fmt.Printf("Erro ao desserializar a mensagem: %v\n", err)
			}
			continue
		}

		rollbackTicket(*result.EventId)
		if err := queue.DeleteMessage(msg.ReceiptHandle); err != nil {
			logger.Error("Ocurred errro when try delete message queue", err)
		}
	}
}
