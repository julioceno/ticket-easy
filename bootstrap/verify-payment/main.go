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

func handler(context context.Context, ticketResponse RequestBody) (events.APIGatewayProxyResponse, error) {
	url := fmt.Sprintf("http://host.docker.internal:8082/tickets/%s/expire-ticket", ticketResponse.TicketId)
	req, err := http.NewRequest("PATCH", url, nil)
	if err != nil {
		log.Fatalf("Occurred error when create request to service", err)
	}

	req.Header.Set("x-api-key", "secret")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	if _, err := client.Do(req); err != nil {
		log.Fatalf("Occurred error when send request to url", err)
	}

	fmt.Println("Completing request...")
	// Responde com sucesso
	response := events.APIGatewayProxyResponse{
		StatusCode: 200,
	}

	return response, nil
}
