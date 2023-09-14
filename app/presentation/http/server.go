package http

import (
	"fmt"
	"log"

	"github.com/DMhattos/FreeCatalogo/config"
	"github.com/gin-gonic/gin"
)

func StartServer(port int) {
	err := config.LoadConfig("config/dev/config.yaml")
	if err != nil {
		log.Fatal()
	}
	// Crie um servidor Gin.
	r := gin.Default()

	// Inicialize o servidor.
	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("Servidor HTTP rodando em http://localhost%s\n", addr)
	err = r.Run(addr)
	if err != nil {
		panic(err)
	}
}
