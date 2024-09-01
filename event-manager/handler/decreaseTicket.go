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

func decreaseTicket(event *schemas.Event) {
	eventId := event.Id.Hex()

	mutexInterface, _ := eventMutexes.LoadOrStore(eventId, &sync.Mutex{})
	mutex := mutexInterface.(*sync.Mutex)
	mutex.Lock()

	ctxMongo, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer func() {
		cancel()
		mutex.Unlock()
		eventMutexes.Delete(eventId)
	}()

	currentEvent := eventsRepository.FindById(&eventId, &ctxMongo)
	notHasTickets := currentEvent.QuantityTickets < 1
	if notHasTickets {
		// TODO: enviar mensagem para fila informando que nao tem mais ingressos
		return
	}

	currentEvent.QuantityTickets--
	eventsRepository.UpdateById(&eventId, &ctxMongo, currentEvent)

	// TODO: enviar mensagem para fila
}

func notifyTicketManager(ticketId *string, message *string) {
	eventUrl := os.Getenv("EVENT_URL")
	apiKey := os.Getenv("EVENT_API_KEY")
	url := fmt.Sprintf("%s/events/%s", eventUrl, *ticketId)

	body := map[string]string{"message": *message}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		logger.Error("Occurred error in marshal message to JSON", err)
		return
	}

	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		logger.Error("Occurred error in build request", err)

		return
	}

	req.Header.Set("x-api-key", apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		logger.Error("Occurred error in call event system", err)
		return
	}

	if ocurredAnyError := response.StatusCode != http.StatusNoContent; ocurredAnyError {
		return
	}

}
