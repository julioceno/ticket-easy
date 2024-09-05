package schemas

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TicketStatus string

const (
	StatusProcessing TicketStatus = "PROCESSING"
	StatusBuying     TicketStatus = "BUYING"
	StatusError      TicketStatus = "ERROR"
	StatusCompleted  TicketStatus = "COMPLETED"
)

type Ticket struct {
	Id           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Status       TicketStatus       `json:"status" bson:"status"`
	Key          string             `json:"key" bson:"key"`
	UserId       string             `json:"userId" bson:"userId"`
	MessageError string             `json:"messageError" bson:"messageError,omitempty"`

	TicketPrice float64   `json:"ticketPrice" bson:"ticketPrice,omitempty"`
	DayEvent    time.Time `json:"dayEvent" bson:"dayEvent,omitempty"`
	EventName   string    `json:"eventName" bson:"eventName,omitempty"`
	EventId     string    `json:"eventId" bson:"eventId"`

	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
}

type TicketResponse struct {
	Id           primitive.ObjectID `json:"_id,omitempty"`
	Status       *TicketStatus      `json:"status,omitempty"`
	UserId       *string            `json:"userId,omitempty"`
	MessageError *string            `json:"messageError,omitempty"`

	TicketPrice *float64   `json:"ticketPrice,omitempty"`
	DayEvent    *time.Time `json:"dayEvent,omitempty"`
	EventName   *string    `json:"eventName,omitempty"`
	EventId     *string    `json:"eventId,omitempty"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (t *Ticket) ToResponse() TicketResponse {
	return TicketResponse{
		Id:           t.Id,
		Status:       &t.Status,
		UserId:       &t.UserId,
		MessageError: &t.MessageError,

		TicketPrice: &t.TicketPrice,
		DayEvent:    &t.DayEvent,
		EventName:   &t.EventName,
		EventId:     &t.EventId,

		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	}
}
