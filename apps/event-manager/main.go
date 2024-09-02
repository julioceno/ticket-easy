package main

import (
	"github.com/julioceno/ticket-easy/apps/event-manager/config/mongoConnection"
	"github.com/julioceno/ticket-easy/apps/event-manager/router"
)

func main() {
	mongoConnection.Seed()
	router.Initialize()
}
