package main

import (
	"github.com/julioceno/ticket-easy/ticket-manager/config/mongoConnection"
	"github.com/julioceno/ticket-easy/ticket-manager/router"
)

func main() {
	mongoConnection.Seed()
	router.Initialize()
}
