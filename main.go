// main.go

package main

import (
	"github.com/phcarneirobc/gopportunities/config"
	"github.com/phcarneirobc/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
    logger = config.GetLogger("main")
    err := config.Init()
    if err != nil {
        logger.Errorf("config initialization error: %v", err)
        return
    }

    // Inicialize o roteador
    routerEngine := router.Initialize()

    // Inicie o servidor Gin
    err = routerEngine.Run(":8080")
    if err != nil {
        logger.Errorf("error starting the server: %v", err)
    }
}
