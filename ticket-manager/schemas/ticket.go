package schemas

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TicketStatus string

const (
	StatusBuy TicketStatus = "BUY"
)

// TODO: salvar dia do evento, salvar dados do evento como o nome e o valor do ingresso
type Ticket struct {
	Id      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Status  TicketStatus       `json:"status" bson:"status"`
	Key     string             `json:"key" bson:"key"`
	EventId string             `json:"eventId" bson:"eventId"`
	UserId  string             `json:"userId" bson:"userId"`

	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
}

type TicketResponse struct {
	Id        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Status    TicketStatus       `json:"status" bson:"status"`
	EventId   string             `json:"eventId" bson:"eventId"`
	UserId    string             `json:"userId" bson:"userId"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
}

func (t *Ticket) ToResponse() TicketResponse {
	return TicketResponse{
		Id:        t.Id,
		Status:    t.Status,
		EventId:   t.EventId,
		UserId:    t.UserId,
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	}
}
