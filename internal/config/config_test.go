package config

import (
	"os"
	"testing"
)

func TestGetEnv(t *testing.T) {
	t.Run("returns environment variable when set", func(t *testing.T) {
		key := "TEST_ENV_VAR"
		expectedValue := "test_value"
		os.Setenv(key, expectedValue)
		defer os.Unsetenv(key)

		result := getEnv(key, "default")

		if result != expectedValue {
			t.Errorf("expected %s, got %s", expectedValue, result)
		}
	})

	t.Run("returns default value when environment variable not set", func(t *testing.T) {
		key := "NONEXISTENT_VAR"
		defaultValue := "default_value"
		os.Unsetenv(key)

		result := getEnv(key, defaultValue)

		if result != defaultValue {
			t.Errorf("expected %s, got %s", defaultValue, result)
		}
	})

	t.Run("returns empty string when environment variable is set to empty", func(t *testing.T) {
		key := "EMPTY_VAR"
		os.Setenv(key, "")
		defer os.Unsetenv(key)

		result := getEnv(key, "default")

		if result != "" {
			t.Errorf("expected empty string, got %s", result)
		}
	})
}

func TestNewConfig(t *testing.T) {
	t.Run("creates config with default values", func(t *testing.T) {
		// Clear all relevant env vars
		os.Unsetenv("DB_USER")
		os.Unsetenv("DB_PASSWORD")
		os.Unsetenv("DB_NAME")
		os.Unsetenv("DB_PORT")
		os.Unsetenv("APP_HOST")

		config := NewConfig()

		expectedDbAddr := "postgres://admin:admin@localhost:5432/app?sslmode=disable"
		if config.DbAddr != expectedDbAddr {
			t.Errorf("expected DbAddr to be %s, got %s", expectedDbAddr, config.DbAddr)
		}
	})

	t.Run("creates config with custom environment variables", func(t *testing.T) {
		os.Setenv("DB_USER", "custom_user")
		os.Setenv("DB_PASSWORD", "custom_pass")
		os.Setenv("DB_NAME", "custom_db")
		os.Setenv("DB_PORT", "5433")
		os.Setenv("APP_HOST", "db.example.com")
		defer func() {
			os.Unsetenv("DB_USER")
			os.Unsetenv("DB_PASSWORD")
			os.Unsetenv("DB_NAME")
			os.Unsetenv("DB_PORT")
			os.Unsetenv("APP_HOST")
		}()

		config := NewConfig()

		expectedDbAddr := "postgres://custom_user:custom_pass@db.example.com:5433/custom_db?sslmode=disable"
		if config.DbAddr != expectedDbAddr {
			t.Errorf("expected DbAddr to be %s, got %s", expectedDbAddr, config.DbAddr)
		}
	})

	t.Run("creates config with partial custom environment variables", func(t *testing.T) {
		os.Unsetenv("DB_USER")
		os.Setenv("DB_PASSWORD", "secret123")
		os.Setenv("DB_NAME", "myapp")
		os.Unsetenv("DB_PORT")
		os.Unsetenv("APP_HOST")
		defer func() {
			os.Unsetenv("DB_PASSWORD")
			os.Unsetenv("DB_NAME")
		}()

		config := NewConfig()

		expectedDbAddr := "postgres://admin:secret123@localhost:5432/myapp?sslmode=disable"
		if config.DbAddr != expectedDbAddr {
			t.Errorf("expected DbAddr to be %s, got %s", expectedDbAddr, config.DbAddr)
		}
	})

	t.Run("returns non-nil config", func(t *testing.T) {
		config := NewConfig()

		if config == nil {
			t.Error("expected config to be non-nil")
		}
	})
}
