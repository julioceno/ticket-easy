package handler

import (
	"github.com/julioceno/ticket-easy/event-manager/config/mongoConnection"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	db *mongo.Client
)

func IntializeHandler() {
	db = mongoConnection.InitConnectionDatabase()
}
