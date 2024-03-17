package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Configuration struct {
	Server   ServerConfig
	Postgres PostgresConfig
}
type ServerConfig struct {
	Port string
}

type PostgresConfig struct {
	PostgresqlHost     string
	PostgresqlPort     string
	PostgresqlUser     string
	PostgresqlPassword string
	PostgresqlDbname   string
	PostgresqlSSLMode  bool
	PostgresAdd        string
}

func NewConfig() *Configuration {
	additional := ""
	if os.Getenv("ENV") != "production" {
		additional = "&sslrootcert=ap-southeast-1-bundle.pem&Timezone=UTC"
		if godotenv.Load() != nil {
			fmt.Println("error loading .env file")
		}
	}
	config := Configuration{
		Server: ServerConfig{
			Port: ":8000",
		},
		Postgres: PostgresConfig{
			PostgresqlHost:     os.Getenv("DB_HOST"),
			PostgresqlPort:     os.Getenv("DB_PORT"),
			PostgresqlUser:     os.Getenv("DB_USERNAME"),
			PostgresqlDbname:   os.Getenv("DB_NAME"),
			PostgresqlPassword: os.Getenv("DB_PASSWORD"),
			PostgresqlSSLMode:  false,
			PostgresAdd:        additional,
		},
	}
	return &config
}
