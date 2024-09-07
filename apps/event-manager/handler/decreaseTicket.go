package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/julioceno/ticket-easy/apps/event-manager/config/logger"
	"github.com/julioceno/ticket-easy/apps/event-manager/handler/queue"
	"github.com/julioceno/ticket-easy/apps/event-manager/schemas"
)

type _body struct {
	MessageError *string `json:"messageError,omitempty"`
	Status       *string `json:"status"`
}

var (
	mutexDecreaseTicket sync.Mutex
)

func decreaseTicket(event *schemas.Event, ticketId *string) {
	msg := updateEventDecreaseTicket(event)
	ocurredError := sentNotifyHttp(ticketId, msg)
	if !ocurredError {
		return
	}
	sendMessageQueue(ticketId, msg)
}

func updateEventDecreaseTicket(event *schemas.Event) *string {
	mutexDecreaseTicket.Lock()
	defer mutexDecreaseTicket.Unlock()

	ctxMongo, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	eventId := event.Id.Hex()
	currentEvent := eventsRepository.FindById(&eventId, &ctxMongo)
	notHasTickets := currentEvent.QuantityTickets < 1
	if notHasTickets {
		msg := "Não existe mais ingressos para esse evento"
		return &msg
	}

	currentEvent.QuantityTickets--
	eventsRepository.UpdateById(&eventId, &ctxMongo, currentEvent)

	return nil
}

func sentNotifyHttp(ticketId *string, message *string) bool {
	ticketUrl := os.Getenv("TICKET_URL")
	apiKey := os.Getenv("TICKET_API_KEY")
	url := fmt.Sprintf("%s/tickets/%s", ticketUrl, *ticketId)

	status := "BUYING"
	body := _body{MessageError: message, Status: &status}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		logger.Error("Occurred error in marshal message to JSON", err)
		return true
	}

	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(jsonBody))
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

func sendMessageQueue(ticketId *string, message *string) {
	type _body struct {
		MessageError *string `json:"messageError,omitempty"`
		TicketId     *string `json:"ticketId"`
	}
	body := _body{TicketId: ticketId, MessageError: message}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		logger.Error("Occurred error in marshal message to JSON", err)
		return
	}

	if err := queue.SendMessage(string(jsonBody)); err != nil {
		logger.Error("Ocurred error when try send message to queue", err)
	}
}
