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

	"github.com/julioceno/ticket-easy/event-manager/config/logger"
	"github.com/julioceno/ticket-easy/event-manager/schemas"
)

var eventMutexes = &sync.Map{}

func decreaseTicket(event *schemas.Event, ticketId *string) {
	msg := updateEventDecreaseTicket(event)
	notifyTicketManager(ticketId, msg)
}

func updateEventDecreaseTicket(event *schemas.Event) *string {
	eventId := event.Id.Hex()
	mutexInterface, _ := eventMutexes.LoadOrStore(eventId, &sync.Mutex{})
	mutex := mutexInterface.(*sync.Mutex)
	mutex.Lock()

	ctxMongo, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer func() {
		cancel()
		mutex.Unlock()
		eventMutexes.Delete(eventId)
	}()

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

func notifyTicketManager(ticketId *string, message *string) bool {
	eventUrl := os.Getenv("TICKET_URL")
	apiKey := os.Getenv("TICKET_API_KEY")
	url := fmt.Sprintf("%s/tickets/%s", eventUrl, *ticketId)

	type _body struct {
		MessageError *string `json:"messageError,omitempty"`
	}
	body := _body{MessageError: message}

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
