package utils

import "errors"

var (
	// Config Errors
	ErrReadFileConfig   = errors.New("erro ao ler o arquivo de configuração")
	ErrDecodeFileConfig = errors.New("erro ao decodificar o arquivo de configuração")

	// Database Errors
	ErrNotOpenConn        = errors.New("erro em conectar ao banco de dados")
	ErrQueryFailed        = errors.New("erro na consulta ao banco de dados")
	ErrInsertFailed       = errors.New("erro ao inserir dados no banco de dados")
	ErrUpdateFailed       = errors.New("erro ao atualizar dados no banco de dados")
	ErrDeleteFailed       = errors.New("erro ao excluir dados no banco de dados")
	ErrTransactionFailed  = errors.New("erro ao executar transação no banco de dados")
	ErrRecordNotFound     = errors.New("registro não encontrado no banco de dados")
	ErrDatabaseConnection = errors.New("erro na conexão com o banco de dados")
)
