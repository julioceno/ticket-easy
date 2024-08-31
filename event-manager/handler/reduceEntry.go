package handler

import "github.com/gin-gonic/gin"

type _body struct {
	ticketId string
}

func ReduceEntry(ctx *gin.Context) {
	// pegar o evento
	eventsRepository.FindById()

	// verificar se o evento ainda tem ingressos

	// se tiver ingresso retornar os dados do evento e começar a processar o ingresso
}
