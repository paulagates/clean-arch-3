package configs

import (
	"log"

	"github.com/spf13/viper"
)

type conf struct {
	DBDriver          string `mapstructure:"DB_DRIVER"`
	DBHost            string `mapstructure:"DB_HOST"`
	DBPort            string `mapstructure:"DB_PORT"`
	DBUser            string `mapstructure:"DB_USER"`
	DBPassword        string `mapstructure:"DB_PASSWORD"`
	DBName            string `mapstructure:"DB_NAME"`
	WebServerPort     string `mapstructure:"WEB_SERVER_PORT"`
	GRPCServerPort    string `mapstructure:"GRPC_SERVER_PORT"`
	GraphQLServerPort string `mapstructure:"GRAPHQL_SERVER_PORT"`
	RabbitMQURL       string `mapstructure:"RABBITMQ_URL"`
}

func LoadConfig() (*conf, error) {
	var cfg conf

	viper.SetConfigFile(".env") // Define o arquivo manualmente
	viper.SetConfigType("env")  // Garante que seja tratado como .env
	viper.AutomaticEnv()        // Lê variáveis de ambiente

	err := viper.ReadInConfig() // Lê o arquivo .env
	if err != nil {
		log.Println("Erro ao carregar .env:", err) // Log para depuração
		return nil, err
	}

	err = viper.Unmarshal(&cfg) // Mapeia as variáveis para a struct
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
