package main

import (
	"github.com/julioceno/ticket-easy/event-manager/config/mongoConnection"
	"github.com/julioceno/ticket-easy/event-manager/router"
)

func main() {
	mongoConnection.Seed()
	router.Initialize()
}
