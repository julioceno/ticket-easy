package repository

import (
	"context"
	"errors"
	"time"

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

func (r *TicketsRepository) FindById(id *string, ctxMongo *context.Context) *schemas.Ticket {
	objId, err := _convertToObjectId(id)
	if err != nil {
		return nil
	}

	var ticket schemas.Ticket
	document := r.collection.FindOne(*ctxMongo, bson.M{"_id": objId})
	if err = document.Decode(&ticket); err != nil {
		return nil
	}

	return &ticket
}

func (r *TicketsRepository) Create(ctxMongo *context.Context, ticket *schemas.Ticket) (*schemas.Ticket, error) {
	now := time.Now()
	ticket.CreatedAt = now
	ticket.UpdatedAt = now

	documentCreated, err := r.collection.InsertOne(*ctxMongo, ticket)
	if err != nil {
		logger.Error("Occured error when create document", err)
		return nil, err
	}

	id := documentCreated.InsertedID.(primitive.ObjectID).Hex()
	document := r.FindById(&id, ctxMongo)
	if document == nil {
		errorCreated := errors.New("document not exists")
		logger.Error("Document not exists", errorCreated)
		return nil, errorCreated
	}

	return document, nil
}

func (r *TicketsRepository) Update(id *string, ctxMongo *context.Context, ticket *schemas.Ticket) (*schemas.Ticket, error) {
	currentTicket := r.FindById(id, ctxMongo)
	if currentTicket == nil {
		errorCreated := errors.New("document not exists")
		logger.Error("Document not exists", errorCreated)
		return nil, errorCreated
	}

	ticket.UpdatedAt = time.Now()
	ticket.CreatedAt = currentTicket.CreatedAt
	update := bson.M{
		"$set": ticket,
	}

	_, err := r.collection.UpdateByID(*ctxMongo, currentTicket.Id, update)
	if err != nil {
		logger.Error("Ocurred error when update document", err)
		return nil, err
	}

	documentUpdated := r.FindById(id, ctxMongo)
	return documentUpdated, nil
}

func _convertToObjectId(id *string) (*primitive.ObjectID, error) {
	objId, err := primitive.ObjectIDFromHex(*id)
	if err != nil {
		return nil, err
	}

	return &objId, nil
}
