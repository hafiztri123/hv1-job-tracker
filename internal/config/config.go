package config

import (
	"fmt"
	"log/slog"
	"os"
)

func NewConfig() *Config {
	dbUser := getEnv("DB_USER", "admin")
	dbPass := getEnv("DB_PASSWORD", "admin")
	dbName := getEnv("DB_NAME", "app")
	dbPort := getEnv("DB_PORT", "5432")
	appHost := getEnv("APP_HOST", "localhost")

	pgUrl := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbUser,
		dbPass,
		appHost,
		dbPort,
		dbName,
	)

	return &Config{
		DbAddr: pgUrl,
	}
}

func getEnv(key, defaultValue string) string {
	value, ok := os.LookupEnv(key)

	if !ok {
		slog.Warn("using default value for env", "key", key)
		return defaultValue
	}

	return value

}
