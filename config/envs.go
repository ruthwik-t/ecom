package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort             string
	DbString               string
	JWTSecret              string
	JWTExpirationInSeconds int64
}

var Envs Config

func init() {
	Envs = initConfig()
}

func initConfig() Config {
	godotenv.Load()

	return Config{
		ServerPort:             getEnv("SERVER_PORT", ":8080"),
		DbString:               getEnv("DB_STRING", ""),
		JWTSecret:              getEnv("JWT_SECRET", "not-so-secret-now-is-it?"),
		JWTExpirationInSeconds: getEnvAsInt("JWT_EXPIRATION_IN_SECONDS", 3600*24*7),
	}
}

func getEnv(key string, fallback string) string {
	if value, isPresent := os.LookupEnv(key); isPresent {
		return value
	}
	return fallback
}

func getEnvAsInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}

		return i
	}

	return fallback
}
