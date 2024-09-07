package handler

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/julioceno/ticket-easy/apps/event-manager/config/logger"
	"github.com/julioceno/ticket-easy/apps/event-manager/schemas"
	"github.com/julioceno/ticket-easy/apps/event-manager/utils"
)

var (
	eventMutexesRollbackTicket = &sync.Map{}
)

func rollbackTicket(eventId string) (*schemas.Event, *utils.ErrorPattern) {
	mutexInterface, _ := eventMutexesRollbackTicket.LoadOrStore(eventId, &sync.Mutex{})
	mutex := mutexInterface.(*sync.Mutex)
	mutex.Lock()

	ctxMongo, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	currentEvent := eventsRepository.FindById(&eventId, &ctxMongo)
	if currentEvent == nil {
		msg := fmt.Sprintf("Event with id %s not exists", currentEvent.Id.Hex())
		logger.Error(msg, nil)
		fmt.Print("Break 1 \n")
		return nil, &utils.ErrorPattern{
			Code:    http.StatusNotFound,
			Message: msg,
		}
	}

	currentEvent.QuantityTickets++
	event, err := eventsRepository.UpdateById(&eventId, &ctxMongo, currentEvent)
	mutex.Unlock()
	if err != nil {
		msg := "Ocurred error when try update event"
		fmt.Print("Break 2 \n")
		logger.Error(msg, err)
		return nil, &utils.ErrorPattern{
			Code:    http.StatusBadRequest,
			Message: msg,
		}
	}

	return event, nil
}
