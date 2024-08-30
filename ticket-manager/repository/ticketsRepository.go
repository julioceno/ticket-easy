package repository

import (
	"context"
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/julioceno/ticket-easy/ticket-manager/config/logger"
	"github.com/julioceno/ticket-easy/ticket-manager/config/mongoConnection"
	"github.com/julioceno/ticket-easy/ticket-manager/schemas"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TicketsRepository struct {
	collection *mongo.Collection
}

func NewTicketRepository() *TicketsRepository {
	return &TicketsRepository{
		collection: mongoConnection.GetMongoCollection(mongoConnection.Database, mongoConnection.DatabaseName.TICKETS),
	}
}

func (r *TicketsRepository) FindById(id string, ctxMongo context.Context) *schemas.Ticket {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil
	}

	var ticket schemas.Ticket
	document := r.collection.FindOne(ctxMongo, bson.M{"_id": objId})
	err = document.Decode(&ticket)
	if err != nil {
		return nil
	}

	return &ticket
}

func (r *TicketsRepository) Create(ctx *gin.Context, ctxMongo context.Context, ticket *schemas.Ticket) (*schemas.Ticket, error) {
	now := time.Now()
	ticket.CreatedAt = now
	ticket.UpdatedAt = now

	documentCreated, err := r.collection.InsertOne(ctxMongo, ticket)
	if err != nil {
		logger.Error("Occured error when create document", err)
		return nil, err
	}

	id := documentCreated.InsertedID.(primitive.ObjectID).Hex()
	document := r.FindById(id, ctxMongo)
	if document == nil {
		errorCreated := errors.New("Document not exists")
		logger.Error("Document not exists", errorCreated)
		return nil, errorCreated
	}

	return document, nil
}

func (r *TicketsRepository) Update(ctx *gin.Context, ctxMongo context.Context, ticket *schemas.Ticket) {
	ticket.UpdatedAt = time.Now()
	// TODO: criar trava pra nao atualizar o createdAt
}
