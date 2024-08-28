package handler

import (
	"github.com/julioceno/ticket-easy/event-manager/repository"
)

var (
	eventsRepository *repository.EventsRepository
)

func init() {
	eventsRepository = repository.NewEventRepository()
}
