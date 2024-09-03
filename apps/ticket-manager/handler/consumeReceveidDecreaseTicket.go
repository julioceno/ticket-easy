package handler

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/julioceno/ticket-easy/apps/ticket-manager/config/logger"
	"github.com/julioceno/ticket-easy/apps/ticket-manager/handler/queue"
)

type ticketAction struct {
	TicketId     *string `json:"ticketId"`
	MessageError *string `json:"messageError"`
}

func startConsumerDecreaseTicket() {
	for {
		msgResult, err := queue.ReceveidMessage()
		if err != nil {
			logger.Error("Ocurred error when trying to get queue messages", err)
			time.Sleep(5 * time.Second)
			continue
		}

		consumeReceveidDecreaseTicket(msgResult)
	}
}

func consumeReceveidDecreaseTicket(msgResult *sqs.ReceiveMessageOutput) {
	messages := msgResult.Messages
	logger.Info(fmt.Sprintf("consuming %v messages", len(messages)))
	for _, msg := range messages {
		var result ticketAction

		if err := json.Unmarshal([]byte(*msg.Body), &result); err != nil {
			fmt.Println("Erro na mensagem:", *msg.Body)
			if e, ok := err.(*json.SyntaxError); ok {
				fmt.Printf("Erro de sintaxe no byte offset %d: %v\n", e.Offset, err)
			} else {
				fmt.Printf("Erro ao desserializar a mensagem: %v\n", err)
			}
			continue
		}

		if err := queue.DeleteMessage(msg.ReceiptHandle); err != nil {
			logger.Error("Ocurred errro when try delete message queue", err)
		}
	}

}
