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
	mutexRollbackTicket sync.Mutex
)

func rollbackTicket(eventId string) (*schemas.Event, *utils.ErrorPattern) {
	mutexRollbackTicket.Lock()
	defer mutexRollbackTicket.Unlock()

	ctxMongo, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	currentEvent := eventsRepository.FindById(&eventId, &ctxMongo)
	if currentEvent == nil {
		msg := fmt.Sprintf("Event with id %s not exists", currentEvent.Id.Hex())
		logger.Error(msg, nil)
		return nil, &utils.ErrorPattern{
			Code:    http.StatusNotFound,
			Message: msg,
		}
	}

	currentEvent.QuantityTickets++
	event, err := eventsRepository.UpdateById(&eventId, &ctxMongo, currentEvent)
	if err != nil {
		msg := "Ocurred error when try update event"
		logger.Error(msg, err)
		return nil, &utils.ErrorPattern{
			Code:    http.StatusBadRequest,
			Message: msg,
		}
	}

	return event, nil
}
