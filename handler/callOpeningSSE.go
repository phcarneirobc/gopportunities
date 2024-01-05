// handler/callOpeningSSEHandler.go (crie um novo arquivo dentro do pacote handler)

package handler

import (
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/phcarneirobc/gopportunities/schemas"
)

var (
	openingsChannel = make(chan schemas.Opening)
	subscribers     []chan<- schemas.Opening
	subscribersLock sync.Mutex
)

// SubscribeOpeningsHandler inscreve um cliente SSE para receber notificações de chamadas de aberturas
func SubscribeOpeningsHandler(ctx *gin.Context) {
	// Configura os cabeçalhos para Server-Sent Events
	ctx.Header("Content-Type", "text/event-stream")
	ctx.Header("Cache-Control", "no-cache")
	ctx.Header("Connection", "keep-alive")
	ctx.Header("Access-Control-Allow-Origin", "*")

	// Canal para receber notificações
	clientChannel := make(chan schemas.Opening)

	// Adiciona o canal do cliente à lista de inscritos
	subscribersLock.Lock()
	subscribers = append(subscribers, clientChannel)
	subscribersLock.Unlock()

	// Defer para remover o canal do cliente quando a conexão for encerrada
	defer func() {
		subscribersLock.Lock()
		defer subscribersLock.Unlock()
		for i, ch := range subscribers {
			if ch == clientChannel {
				subscribers = append(subscribers[:i], subscribers[i+1:]...)
				close(clientChannel)
				break
			}
		}
	}()

	// Loop para enviar notificações ao cliente
	for opening := range clientChannel {
		ctx.SSEvent("callOpening", opening)
		ctx.Writer.Flush()
	}
}

// NotifyOpeningCalled notifica todos os inscritos que uma abertura foi chamada
func NotifyOpeningCalled(opening schemas.Opening) {
	subscribersLock.Lock()
	defer subscribersLock.Unlock()

	for _, ch := range subscribers {
		ch <- opening
	}
}
