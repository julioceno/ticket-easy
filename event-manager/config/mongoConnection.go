package config

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitConnectionDatabase() *mongo.Client {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017")) // TODO: adicionar numa variavel de ambiente

	if err != nil {
		panic("Occurred an error with mongo connection")
	}

	if err := client.Ping(ctx, options.Client().ReadPreference); err != nil {
		panic("Occured an error in make ping in mongo connection")
	}

	defer func() {
		fmt.Print("Disconectin database...")
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	return client
}

func GetMongoCollection(db *mongo.Client, nameCollection string) *mongo.Collection {
	eventsCollection := db.Database("events").Collection(nameCollection)
	return eventsCollection
}
