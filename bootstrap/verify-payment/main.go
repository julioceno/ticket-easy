package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type RequestBody struct {
	TicketId string `json:"ticketId"`
}

func main() {
	lambda.Start(handler)
}

func handler(ctx context.Context, ticketResponse RequestBody) (events.APIGatewayProxyResponse, error) {
	url := fmt.Sprintf("http://host.docker.internal:8082/tickets/%s/expire-ticket", ticketResponse.TicketId)
	req, err := http.NewRequest("PATCH", url, nil)
	if err != nil {
		log.Printf("Error creating request: %v", err)
	}

	req.Header.Set("x-api-key", "secret")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	go client.Do(req)

	log.Printf("Completed request for ticket %s...", ticketResponse.TicketId)
	return events.APIGatewayProxyResponse{StatusCode: 200}, nil
}
