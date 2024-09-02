package schemas

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Event struct {
	Id              primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name            string             `json:"name" bson:"name"`
	Description     string             `json:"description" bson:"description"`
	TicketValue     float64            `json:"ticketValue" bson:"ticketValue"`
	ImagesUrl       []string           `json:"imagesUrl" bson:"imagesUrl"`
	QuantityTickets int                `json:"quantityTickets" bson:"quantityTickets"`
	OccuredAt       time.Time          `json:"occuredAt" bson:"occuredAt"`
}
