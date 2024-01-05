package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Initialize() *gin.Engine {
	// Inicialize o roteador
	router := gin.Default()

	// Configuração básica do CORS (permitindo todas as origens)
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"} // ou ajuste para a origem específica do seu frontend
	router.Use(cors.New(config))

	// Inicialize Rotas
	initializeRoutes(router)

	// Retorne o roteador
	return router
}