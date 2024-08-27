package schemas

import "time"

type Event struct {
	Name            string    `bson:"name"`
	Description     string    `bson:"description"`
	TicketValue     float64   `bson:"ticketValue"`
	ImagesUrl       []string  `bson:"imagesUrl"`
	QuantityTickets int       `bson:"quantityTickets"`
	OccuredAt       time.Time `bson:"occuredAt"`
}
