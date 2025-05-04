package bootstrap

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Env struct {
	DBHost                 string
	DBPort                 string
	DBUser                 string
	DBPassword             string
	DBName                 string
	AppPort                string
	BotToken               string
	AccessTokenSecret      string
	RefreshTokenSecret     string
	AccessTokenExpiryHour  int
	RefreshTokenExpiryHour int
}

func NewEnv() *Env {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file : ", err)
	}


	return &Env{
		DBHost:                 os.Getenv("DB_HOST"),
		DBPort:                 os.Getenv("DB_PORT"),
		DBUser:                 os.Getenv("DB_USER"),
		DBPassword:             os.Getenv("DB_PASSWORD"),
		DBName:                 os.Getenv("DB_NAME"),
		AppPort:                os.Getenv("APP_PORT"),
		BotToken:               os.Getenv("BOT_TOKEN"),
		AccessTokenSecret:      os.Getenv("ACCESS_TOKEN_SECRET"),
		RefreshTokenSecret:     os.Getenv("REFRESH_TOKEN_SECRET"),
		AccessTokenExpiryHour:  getEnvAsInt("ACCESS_TOKEN_EXPIRY_HOUR"),
		RefreshTokenExpiryHour: getEnvAsInt("REFRESH_TOKEN_EXPIRY_HOUR"),
	}
}

func getEnvAsInt(name string) int {
    valueStr := os.Getenv(name)
    value, err := strconv.Atoi(valueStr)

	if err != nil {
		log.Fatal("Error loading .env file : ", err)
	}

    return value
}
