package schemas

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Event struct {
	Id              primitive.ObjectID `bson:"_id,omitempty"`
	Name            string             `bson:"name"`
	Description     string             `bson:"description"`
	TicketValue     float64            `bson:"ticketValue"`
	ImagesUrl       []string           `bson:"imagesUrl"`
	QuantityTickets int                `bson:"quantityTickets"`
	OccuredAt       time.Time          `bson:"occuredAt"`
}
