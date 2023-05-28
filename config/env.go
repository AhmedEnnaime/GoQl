package config

import (
	"log"

	"github.com/caarlos0/env/v8"
	"github.com/joho/godotenv"
)

type dbEnvVariables struct {
	DbUser     string `env:"POSTGRES_USER"`
	DbHost     string `env:"DB_HOST"`
	DbPassword string `env:"POSTGRES_PASSWORD"`
	DbName     string `env:"POSTGRES_DB"`
	DbPort     string `env:"DB_PORT"`
}

var DbEnvConfig dbEnvVariables

func LoadEnvVariables() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if err := env.Parse(&DbEnvConfig); err != nil {
		log.Fatal("Error parsing .env file")
	}
}
