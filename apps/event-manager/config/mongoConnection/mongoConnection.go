package mongoConnection

import (
	"context"
	"os"

	"github.com/joho/godotenv"
	"github.com/julioceno/ticket-easy/apps/event-manager/config/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Database *mongo.Client
)

func init() {
	if err := godotenv.Load(); err != nil {
		logger.Fatal("Error loading .env file", err)
	}

	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		logger.Fatal("Database url not exists", nil)
	}

	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbUrl))

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
