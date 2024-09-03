package handler

import "github.com/julioceno/ticket-easy/apps/ticket-manager/repository"

var (
	ticketsRepository *repository.TicketsRepository
)

func init() {
	ticketsRepository = repository.NewTicketRepository()
	go startConsumerDecreaseTicket()
}
