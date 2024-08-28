package mongoConnection

import (
	"context"

	"github.com/julioceno/ticket-easy/event-manager/config/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Database *mongo.Client
)

func init() {
	ctx := context.Background()
	// TODO: adicionar numa variavel de ambiente
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		logger.Fatal("Occurred an error with mongo connection", err)
	}

	if err := client.Ping(ctx, options.Client().ReadPreference); err != nil {
		logger.Fatal("Occured an error in make ping in mongo connection", err)
	}

	Database = client
}

func GetMongoCollection(db *mongo.Client, nameCollection string) *mongo.Collection {
	eventsCollection := db.Database("events").Collection(nameCollection)
	return eventsCollection
}
