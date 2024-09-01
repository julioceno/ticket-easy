package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/julioceno/ticket-easy/ticket-manager/config/logger"
	"github.com/julioceno/ticket-easy/ticket-manager/schemas"
	"github.com/julioceno/ticket-easy/ticket-manager/utils"
)

type _CreateTicketbody struct {
	UserId  string `json:"userId" validate:"required"`
	EventId string `json:"eventId" validate:"required"`
}

type _eventData struct {
	ID              string    `json:"_id"`
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	TicketValue     float64   `json:"ticketValue"`
	ImagesUrl       []string  `json:"imagesUrl"`
	QuantityTickets int       `json:"quantityTickets"`
	OccuredAt       time.Time `json:"occuredAt"`
}

type _responseEvent struct {
	Data    _eventData `json:"data"`
	Message string     `json:"message"`
	Status  int        `json:"status"`
}

func CreateTicket(ctx *gin.Context) {
	var body _CreateTicketbody
	if err := utils.DecodeBody(ctx, &body); err != nil {
		logger.Error("Ocurred error when try decode body", err)
		utils.SendError(ctx, http.StatusBadRequest, "Ocorreu um erro ao tentar criar o ticket")
		return
	}

	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		logger.Error("Ocurred error validate body", err)
		utils.SendError(ctx, http.StatusBadRequest, "O body da requisição está incorreto")
		return
	}

	ticket, err := createTicket(body)
	if err != nil {
		logger.Error("Ocurred error when format body to create ticket", err)
		utils.SendError(ctx, http.StatusBadRequest, "Ocorreu um erro ao tentar criar o ticket")
		return
	}

	ctxMongo, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ticketCreated, err := ticketsRepository.Create(&ctxMongo, ticket)
	if err != nil {
		logger.Error("Ocurred error when try create ticket", err)
		utils.SendError(ctx, http.StatusBadRequest, "Ocorreu um erro ao tentar criar o ticket")
		return
	}

	responseEvent, messageError := buyTicket(ctx, ticketCreated)
	if messageError != nil {
		updateEventError(&ctxMongo, messageError, ticketCreated)
		utils.SendError(ctx, http.StatusBadGateway, *messageError)
		return
	}

	updateEventInfo(&ctxMongo, responseEvent, ticketCreated)

	response := ticketCreated.ToResponse()
	responseStatus := http.StatusCreated
	utils.SendSuccess(utils.SendSuccesStruct{ctx, "POST", response, &responseStatus})
}

func createTicket(body _CreateTicketbody) (*schemas.Ticket, error) {
	ticket := schemas.Ticket{
		Status:  schemas.StatusProcessing,
		EventId: body.EventId,
		UserId:  body.UserId,
	}

	return &ticket, nil
}

func buyTicket(ctx *gin.Context, ticket *schemas.Ticket) (*_responseEvent, *string) {
	messageError := "Ocorreu um erro ao efetuar a compra, entre em contato com nossa equipe de suporte"
	eventUrl := os.Getenv("EVENT_URL")
	apiKey := os.Getenv("EVENT_API_KEY")
	url := fmt.Sprintf("%s/events/%s/reduce-ticket", eventUrl, ticket.EventId)

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		logger.Error("Occurred error in build request", err)
		return nil, &messageError
	}
	req.Header.Set("x-api-key", apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		logger.Error("Occurred error in call event system", err)
		return nil, &messageError
	}

	if response.StatusCode != http.StatusOK {
		err := errors.New("request is fail")
		logger.Error(fmt.Sprintf("request can not process, status code %v", response.StatusCode), err)
		return nil, &messageError
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		logger.Error("Occurred error when try read body", err)
		utils.SendError(ctx, http.StatusBadGateway, messageError)
		return nil, &messageError
	}

	var responseStruct _responseEvent
	err = json.Unmarshal(body, &responseStruct)

	if err != nil {
		logger.Error("Occurred error when try unmarshal body", err)
		utils.SendError(ctx, http.StatusBadGateway, "Ocorreu um problema ao processar a resposta do servidor")
		return nil, &messageError
	}

	return &responseStruct, nil
}

func updateEventInfo(ctxMongo *context.Context, responseEvent *_responseEvent, ticket *schemas.Ticket) {
	ticket.TicketPrice = responseEvent.Data.TicketValue
	ticket.EventName = responseEvent.Data.Name
	ticket.DayEvent = responseEvent.Data.OccuredAt

	id := ticket.Id.Hex()
	ticketsRepository.Update(&id, ctxMongo, ticket)
}
