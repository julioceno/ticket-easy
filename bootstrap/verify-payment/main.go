package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

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
		return events.APIGatewayProxyResponse{StatusCode: 500}, err
	}

	req.Header.Set("x-api-key", "secret")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error sending request: %v", err)
		return events.APIGatewayProxyResponse{StatusCode: 500}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Service returned non-OK status: %d", resp.StatusCode)
		return events.APIGatewayProxyResponse{StatusCode: resp.StatusCode}, nil
	}

	log.Println("Completing request...")
	return events.APIGatewayProxyResponse{StatusCode: 200}, nil
}
