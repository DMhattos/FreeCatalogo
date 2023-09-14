package utils

import "errors"

var (
	//Config Erros
	ErrReadFileConfig   = errors.New("erro ao ler o arquivo de configuração")
	ErrDecodeFileConfig = errors.New("erro ao decodificar o arquivo de configuração")
)
