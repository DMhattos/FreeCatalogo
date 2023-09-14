package http

import (
	"fmt"
	"log"

	"github.com/DMhattos/FreeCatalogo/app/application/usecase"
	"github.com/DMhattos/FreeCatalogo/app/infra/database"
	"github.com/DMhattos/FreeCatalogo/app/presentation/http/handlers"
	"github.com/DMhattos/FreeCatalogo/app/presentation/http/routes"
	"github.com/DMhattos/FreeCatalogo/config"
	"github.com/gin-gonic/gin"
)

// StartServer inicia o servidor HTTP na porta especificada.
// O parâmetro "port" determina a porta na qual o servidor será executado.
func StartServer(port int) {
	err := config.LoadConfig("config/dev/config.yaml")
	if err != nil {
		log.Fatal("Erro ao carregar a configuração:", err)
	}

	// Crie um servidor Gin com as configurações padrão.
	r := gin.Default()

	// Abra uma conexão com o banco de dados.
	db, err := database.OpenConnection()
	if err != nil {
		log.Fatal("Erro ao abrir a conexão com o banco de dados:", err)
	}
	defer db.Close()
	// Crie instâncias dos repositórios usando a conexão com o banco de dados.
	repositories := database.NewRepositories(db)

	// Crie instâncias dos casos de uso usando os repositórios.
	usecases := usecase.NewUsecases(repositories)

	// Crie instâncias dos handlers da aplicação usando os casos de uso.
	handlers := handlers.NewHandlers(usecases)

	// Configure as rotas da aplicação com base nos handlers.
	routes.SetupRoutes(handlers, r)

	// Inicialize o servidor na porta especificada.
	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("Servidor HTTP rodando em http://localhost%s\n", addr)
	err = r.Run(addr)
	if err != nil {
		panic(err)
	}
}
