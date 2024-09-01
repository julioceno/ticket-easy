package handler

import (
	"context"
	"sync"
	"time"

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
	if currentEvent.QuantityTickets <= 0 {
		// TODO: enviar mensagem para fila informando que nao tem mais ingressos
		return
	}

	currentEvent.QuantityTickets--
	eventsRepository.UpdateById(&eventId, &ctxMongo, currentEvent)

	// TODO: enviar mensagem para fila
}
