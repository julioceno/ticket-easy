package handler

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/julioceno/ticket-easy/event-manager/schemas"
	"github.com/julioceno/ticket-easy/event-manager/utils"
	"go.mongodb.org/mongo-driver/bson"
)

func GetEvents(ctx *gin.Context) {
	fmt.Println("Getting collection")
	eventsCollection := db.Database("events").Collection("events")

	fmt.Println("Creating context mongo")
	ctxMongo, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	fmt.Println("Creating context mongo", eventsCollection, ctxMongo)
	cursor, err := eventsCollection.Find(ctxMongo, bson.D{})
	if err != nil {
		log.Fatal("Occurred an error", err)
	}

	defer cursor.Close(ctx)

	events := []schemas.Event{}

	fmt.Println("Calling cursor")
	for cursor.Next(ctxMongo) {
		fmt.Println("Decoding event...")
		var event schemas.Event
		err := cursor.Decode(&event)
		if err != nil {
			log.Println("Error decoding event:", err)
			ctx.AbortWithStatusJSON(500, gin.H{"error": "Error decoding event"})
			return
		}
		events = append(events, event)
	}

	utils.SendSuccess(ctx, "GET", events)
}
