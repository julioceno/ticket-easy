package repository

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/julioceno/ticket-easy/event-manager/config/logger"
	"github.com/julioceno/ticket-easy/event-manager/config/mongoConnection"
	"github.com/julioceno/ticket-easy/event-manager/schemas"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type EventsRepository struct {
	collection *mongo.Collection
}

func NewEventRepository() *EventsRepository {
	return &EventsRepository{
		collection: mongoConnection.GetMongoCollection(mongoConnection.Database, mongoConnection.DatabaseName.EVENTS),
	}
}

func (r *EventsRepository) FetchEvents(ctx *gin.Context, ctxMongo context.Context) ([]schemas.Event, error) {
	filter := createFilter(ctx)
	opts, _ := createSortOptions(ctx)

	cursor, err := r.collection.Find(ctxMongo, filter, opts)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctxMongo)

	var events []schemas.Event
	for cursor.Next(ctxMongo) {
		var event schemas.Event
		if err := cursor.Decode(&event); err != nil {
			logger.Error("Error decoding event:", err)
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

func (r *EventsRepository) CountEvents(ctx *gin.Context, ctxMongo context.Context) (int64, error) {
	filter := createFilter(ctx)
	count, err := r.collection.CountDocuments(ctxMongo, filter)
	if err != nil {
		logger.Error("Error counting events:", err)
		return 0, err
	}
	return count, nil
}

func (r *EventsRepository) FindById(id string, ctxMongo context.Context) *schemas.Event {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil
	}

	var event schemas.Event
	document := r.collection.FindOne(ctxMongo, bson.M{"_id": objID})
	err = document.Decode(&event)
	if err != nil {
		return nil
	}

	return &event
}

func createFilter(ctx *gin.Context) primitive.D {
	name := ctx.Query("name")
	filter := bson.D{{"name", primitive.Regex{Pattern: name, Options: "i"}}}

	return filter
}

func createSortOptions(ctx *gin.Context) (*options.FindOptions, error) {
	skip := ctx.DefaultQuery("skip", "0")
	limit := ctx.DefaultQuery("limit", "10")

	skipBase64, err := convertToBase64(skip)
	if err != nil {
		return nil, err
	}

	limitBase64, err := convertToBase64(limit)
	if err != nil {
		return nil, err
	}

	opts := options.Find().SetSkip(skipBase64).SetLimit(limitBase64)
	return opts, nil
}

func convertToBase64(value string) (int64, error) {
	valueNumeric, err := strconv.Atoi(value)
	if err != nil {
		return 0, err
	}

	return int64(valueNumeric), nil
}
