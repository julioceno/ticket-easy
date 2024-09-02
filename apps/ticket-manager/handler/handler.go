package handler

import "github.com/julioceno/ticket-easy/ticket-manager/repository"

var (
	ticketsRepository *repository.TicketsRepository
)

func init() {
	ticketsRepository = repository.NewTicketRepository()
}
