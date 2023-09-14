// Pacote config é responsável por carregar e gerenciar configurações de aplicativo, como configurações de banco de dados.
package config

import (
	"github.com/DMhattos/FreeCatalogo/app/utils"
	"github.com/spf13/viper"
)

var data *Config

// Config é a estrutura principal que armazena as configurações do aplicativo.
type Config struct {
	Database ConfigDatabase `mapstructure:"database"`
}

// ConfigDatabase é a estrutura que armazena as configurações específicas do banco de dados.
type ConfigDatabase struct {
	Host     string `mapstructure:"host"`     // Host do banco de dados.
	Port     uint   `mapstructure:"port"`     // Porta do banco de dados.
	DbName   string `mapstructure:"dbName"`   // Nome do banco de dados.
	User     string `mapstructure:"user"`     // Nome de usuário do banco de dados.
	Password string `mapstructure:"password"` // Senha do banco de dados.
}

// LoadConfig carrega as configurações do arquivo especificado.
//
// Ele configura o Viper para ler o arquivo de configuração em pathFile,
// decodifica as configurações e as armazena na estrutura de dados Config.
// Retorna um erro se houver problemas ao ler ou decodificar o arquivo.
func LoadConfig(pathFile string) error {
	viper.SetConfigFile(pathFile)

	if err := viper.ReadInConfig(); err != nil {
		return utils.ErrReadFileConfig
	}

	if err := viper.Unmarshal(&data); err != nil {
		return utils.ErrDecodeFileConfig
	}

	return nil
}

// GetConfig retorna a estrutura de configuração carregada.
//
// Permite que outras partes do aplicativo acessem as configurações carregadas
// para uso em todo o aplicativo.
func GetConfig() *Config {
	return data
}
