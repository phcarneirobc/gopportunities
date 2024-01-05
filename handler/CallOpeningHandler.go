// handler/callOpeningHandler.go

package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/phcarneirobc/gopportunities/schemas"
)

// CallOpeningHandler manipula a chamada de uma abertura
func CallOpeningHandler(ctx *gin.Context) {
	var request schemas.CallOpeningRequest

	if err := ctx.BindJSON(&request); err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := request.Validate(); err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening, err := callOpeningLogic(request.ID) // Implemente sua lógica real de chamada de abertura aqui

	if err != nil {
		sendError(ctx, http.StatusInternalServerError, "Erro ao chamar abertura")
		return
	}

	// Notifica os inscritos sobre a abertura chamada
	NotifyOpeningCalled(opening)

	sendSucess(ctx, "call-opening", opening)
}

// Adicione a sua lógica real de chamada de abertura aqui
func callOpeningLogic(openingID uint) (schemas.Opening, error) {
	// Implemente sua lógica real de chamada de abertura aqui
	// Por exemplo, consulte a abertura no banco de dados e marque como chamada
	return schemas.Opening{}, nil
}
