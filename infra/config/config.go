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
		// ISSUE: .env load failure is non-fatal. In production there is no .env file,
		// so this always prints, but the app continues. All config depends on env vars
		// being present, so this creates a confusing failure mode: app starts but
		// panics later with an opaque database connection error. Should validate that
		// required env vars are present, or treat a missing .env as a fatal error here.
		fmt.Println("Error loading .env file")
	}
}

func GetConfig() appConfig {
	// ISSUE: No validation that required env vars (DB_HOST, DB_USER, DB_PASSWORD,
	// DB_NAME, DB_PORT) are non-empty. If any is missing, GetConfig returns an empty
	// string, producing a malformed DSN and an opaque connection error downstream.
	// ISSUE: JWT_SECRET_KEY is not validated. If missing, the app uses an empty HMAC
	// key — tokens signed with an empty key can be trivially forged. This is a critical
	// security vulnerability.
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
	// ISSUE: No DefaultValue (and no PORT validation). If PORT env var is missing,
	// `r.Run(":")` is called with an empty string, failing to listen on a predictable
	// port. Should fall back to a default (e.g. "8080").
	return serverConfig{
		Port: os.Getenv("PORT"),
	}
}
