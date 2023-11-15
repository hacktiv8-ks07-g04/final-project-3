package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type appConfig struct {
	DBHost       string
	DBUser       string
	DBPassword   string
	DBName       string
	DBPort       string
	JWTSecretKey string
}

type serverConfig struct {
	Port string
}

func LoadAppConfig() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}
}

func GetConfig() appConfig {
	return appConfig{
		DBHost:       os.Getenv("DB_HOST"),
		DBPort:       os.Getenv("DB_PORT"),
		DBUser:       os.Getenv("DB_USER"),
		DBPassword:   os.Getenv("DB_PASSWORD"),
		DBName:       os.Getenv("DB_NAME"),
		JWTSecretKey: os.Getenv("JWT_SECRET_KEY"),
	}
}

func Server() serverConfig {
	return serverConfig{
		Port: os.Getenv("PORT"),
	}
}
