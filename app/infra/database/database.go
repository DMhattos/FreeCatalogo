package database

import (
	"database/sql"
	"fmt"

	"github.com/DMhattos/FreeCatalogo/config"
	_ "github.com/lib/pq"
)

// OpenConnection abre uma conexão com o banco de dados PostgreSQL.
func OpenConnection() (*sql.DB, error) {
	// Obtenha as configurações do banco de dados do arquivo de configuração.
	dbConfig := config.GetConfig().Database

	// Valide as configurações.
	if dbConfig.Host == "" || dbConfig.Port == 0 || dbConfig.User == "" ||
		dbConfig.Password == "" || dbConfig.DbName == "" {
		return nil, fmt.Errorf("configurações de banco de dados incompletas")
	}

	// Crie a string de conexão.
	connectionString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.DbName)

	// Abra a conexão com o banco de dados.
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, fmt.Errorf("falha ao abrir a conexão com o banco de dados: %v", err)
	}

	// Verifique se a conexão é válida.
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("falha ao verificar a conexão com o banco de dados: %v", err)
	}

	return db, nil
}

// CloseConnection fecha a conexão com o banco de dados.
func CloseConnection(db *sql.DB) {
	if db != nil {
		db.Close()
	}
}
