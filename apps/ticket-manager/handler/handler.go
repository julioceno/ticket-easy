package handler

import (
	"github.com/julioceno/ticket-easy/apps/ticket-manager/handler/awsServices"
	"github.com/julioceno/ticket-easy/apps/ticket-manager/repository"
)

var (
	ticketsRepository *repository.TicketsRepository
)

func init() {
	awsServices.Initialize()
	ticketsRepository = repository.NewTicketRepository()
	go startConsumerDecreaseTicket()
}
