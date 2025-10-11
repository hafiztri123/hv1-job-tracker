package utils

import (
	"log/slog"
	"os"
)

func GetEnv(key, defaultValue string) string {
	value, ok := os.LookupEnv(key)

	if !ok {
		slog.Warn("using default value for env", "key", key)
		return defaultValue
	}

	return value

}
