package handler

import (
	"context"
	"encoding/json"
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

type _body struct {
	UserId  string `json:"userId" validate:"required"`
	EventId string `json:"eventId" validate:"required"`
}

type _responseEvent struct {
	Id              string    `json:"_id""`
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	TicketValue     float64   `json:"ticketValue"`
	ImagesUrl       []string  `json:"imagesUrl"`
	QuantityTickets int       `json:"quantityTickets"`
	OccuredAt       time.Time `json:"occuredAt"`
}

func CreateTicket(ctx *gin.Context) {
	var body _body
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

	ticketCreated, err := ticketsRepository.Create(ctx, ctxMongo, ticket)
	if err != nil {
		logger.Error("Ocurred error when try create ticket", err)
		utils.SendError(ctx, http.StatusBadRequest, "Ocorreu um erro ao tentar criar o ticket")
		return
	}

	_, messageError := buyTicket(ctx, ticketCreated)
	if messageError != nil {
		// TODO: atualizar ticket setando a mensagem de erro
		return
	}

	response := ticketCreated.ToResponse()
	responseStatus := http.StatusCreated
	utils.SendSuccess(utils.SendSuccesStruct{ctx, "POST", response, &responseStatus})
}

func createTicket(body _body) (*schemas.Ticket, error) {
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
	url := fmt.Sprintf("%s/events/%s/reduce-ticket", eventUrl, ticket.Id)

	resp, err := http.Post(url, "application/json", nil)
	if err != nil {
		logger.Error("Occurred error in call event system", err)
		utils.SendError(ctx, http.StatusBadGateway, "Ocorreu um problema ao efetuar ao solicitar a compra do ingresso")
		return nil, &messageError
	}

	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		logger.Error("Occurred error when try read body", err)
		utils.SendError(ctx, http.StatusBadGateway, "Ocorreu um problema ao efetuar ao solicitar a compra do ingresso")
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
