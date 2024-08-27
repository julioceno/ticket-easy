package mongoConnection

import (
	"context"
	"time"

	"github.com/julioceno/ticket-easy/event-manager/config/logger"
	"github.com/julioceno/ticket-easy/event-manager/schemas"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Seed() {
	logger.Info("Building events...")
	events := buildEvents()
	ctx := context.Background()

	db := InitConnectionDatabase()
	collection := GetMongoCollection(db, "events")

	hasDocuments := verifyIfHasDocuments(collection, ctx)
	if hasDocuments {
		logger.Info("Has documents in event collection")
		return
	}

	logger.Info("Insert events in collection...")
	_, err := collection.InsertMany(ctx, events)
	if err != nil {
		logger.Fatal("Occured error in application seed", err)
	}
}

func verifyIfHasDocuments(collection *mongo.Collection, ctx context.Context) bool {
	logger.Info("Verify if has documents in events collecion...")
	documents, err := collection.CountDocuments(ctx, bson.D{})

	if err != nil {
		logger.Fatal("Ocurred error in get documents from collection", err)
	}

	return documents > 0
}

func buildEvents() []interface{} {
	tomorrow := time.Now().AddDate(0, 0, 1)

	event1 := schemas.Event{
		Name:            "Felipe Ret",
		Description:     "Show do Felipe Ret",
		TicketValue:     50.50,
		QuantityTickets: 1500,
		ImagesUrl: []string{
			"https://www.aen.pr.gov.br/sites/default/arquivos_restritos/files/imagem/2023-02/img_2646_1.jpg",
			"https://www.aen.pr.gov.br/sites/default/arquivos_restritos/files/imagem/2023-02/img_2646_1.jpg",
			"https://www.aen.pr.gov.br/sites/default/arquivos_restritos/files/imagem/2023-02/img_2646_1.jpg",
		},
		OccuredAt: tomorrow,
	}

	event2 := schemas.Event{
		Name:            "Lil Wind",
		Description:     "Show do Lil Wind",
		TicketValue:     45.00,
		QuantityTickets: 1100,
		ImagesUrl: []string{
			"https://www.aen.pr.gov.br/sites/default/arquivos_restritos/files/imagem/2023-02/img_2646_1.jpg",
			"https://www.aen.pr.gov.br/sites/default/arquivos_restritos/files/imagem/2023-02/img_2646_1.jpg",
			"https://www.aen.pr.gov.br/sites/default/arquivos_restritos/files/imagem/2023-02/img_2646_1.jpg",
		},
		OccuredAt: tomorrow,
	}

	event3 := schemas.Event{
		Name:            "Emicida",
		Description:     "Show do Emicida",
		TicketValue:     20,
		QuantityTickets: 1500,
		ImagesUrl: []string{
			"https://www.aen.pr.gov.br/sites/default/arquivos_restritos/files/imagem/2023-02/img_2646_1.jpg",
			"https://www.aen.pr.gov.br/sites/default/arquivos_restritos/files/imagem/2023-02/img_2646_1.jpg",
			"https://www.aen.pr.gov.br/sites/default/arquivos_restritos/files/imagem/2023-02/img_2646_1.jpg",
		},
		OccuredAt: tomorrow,
	}

	event4 := schemas.Event{
		Name:            "Racionais MC's",
		Description:     "Show do Racionais MC's",
		TicketValue:     100,
		QuantityTickets: 1500,
		ImagesUrl: []string{
			"https://www.aen.pr.gov.br/sites/default/arquivos_restritos/files/imagem/2023-02/img_2646_1.jpg",
			"https://www.aen.pr.gov.br/sites/default/arquivos_restritos/files/imagem/2023-02/img_2646_1.jpg",
			"https://www.aen.pr.gov.br/sites/default/arquivos_restritos/files/imagem/2023-02/img_2646_1.jpg",
		},
		OccuredAt: tomorrow,
	}

	events := make([]interface{}, 0)

	events = append(events, event1)
	events = append(events, event2)
	events = append(events, event3)
	events = append(events, event4)

	return events
}
