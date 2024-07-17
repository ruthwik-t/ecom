package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort string
	DbString   string
}

var Envs Config

func init() {
	Envs = initConfig()
}

func initConfig() Config {
	godotenv.Load()

	return Config{
		ServerPort: getEnv("SERVER_PORT", ":8080"),
		DbString:   getEnv("DB_STRING", ""),
	}
}

func getEnv(key string, fallback string) string {
	if value, isPresent := os.LookupEnv(key); isPresent {
		return value
	}
	return fallback
}
