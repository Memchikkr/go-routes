package bootstrap

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	AppPort    string
}

func NewEnv() *Env {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file : ", err)
	}
	return &Env{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		AppPort:    os.Getenv("APP_PORT"),
	}
}
