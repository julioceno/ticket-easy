package schemas

import "time"

type Event struct {
	Name        string        `bson:"name"`
	Description string        `bson:"description"`
	ticketValue float64       `bson:"ticketValue"`
	imagesUrl   []string      `bson:"imagesUrl"`
	occuredAt   time.Duration `bson:"occuredAt"`
}
