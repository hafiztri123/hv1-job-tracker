package config

import (
	"fmt"
	"hafiztri123/hv1-job-tracker/internal/user"
	"hafiztri123/hv1-job-tracker/internal/utils"
	"log/slog"
	"strconv"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewConfig() *Config {

	dbUser := utils.GetEnv("DB_USER", "admin")
	dbPass := utils.GetEnv("DB_PASSWORD", "admin")
	dbName := utils.GetEnv("DB_NAME", "app")
	dbPort := utils.GetEnv("DB_PORT", "5432")
	appHost := utils.GetEnv("APP_HOST", "localhost")

	pgUrl := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbUser,
		dbPass,
		appHost,
		dbPort,
		dbName,
	)

	maxConns := utils.GetEnv("DB_MAX_CONNS", "10")
	maxConnsInt, err := strconv.Atoi(maxConns)

	if err != nil {
		slog.Warn("failed to set max conns, use default value", "error", err)
		maxConnsInt = 10
	}

	return &Config{
		DbAddr:     pgUrl,
		DbMaxConns: int32(maxConnsInt),
	}
}

func NewRepositories(db *pgxpool.Pool) *Repositories {
	return &Repositories{
		UserRepository: user.NewUserRepository(db),
	}
}

func NewService(r *Repositories) *Services {
	return &Services{
		UserService: user.NewUserService(r.UserRepository),
	}
}
