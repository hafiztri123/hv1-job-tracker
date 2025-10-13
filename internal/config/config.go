package config

import (
	"fmt"
	"hafiztri123/hv1-job-tracker/internal/applications"
	"hafiztri123/hv1-job-tracker/internal/middleware"
	"hafiztri123/hv1-job-tracker/internal/user"
	"hafiztri123/hv1-job-tracker/internal/utils"
	"log/slog"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
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

func NewRouterConfig() fiber.Config {
	baseConfig := fiber.Config{
		AppName:      "Job Tracker v1.0",
		ServerHeader: "Fiber",
		ErrorHandler: middleware.ErrorHandler(),
		BodyLimit:    10 * 10 * 1024,
	}

	isDev, err := strconv.ParseBool(utils.GetEnv("IS_DEV", "true"))
	if err != nil {
		isDev = true
	}

	if isDev {
		slog.Info("using router config", "mode", "development")
		baseConfig.ReadTimeout = 300 * time.Second
		baseConfig.WriteTimeout = 60 * time.Second
		baseConfig.IdleTimeout = 600 * time.Second
		baseConfig.Prefork = false

		return baseConfig
	} else {
		slog.Info("using router config", "mode", "production")
		baseConfig.ReadTimeout = 120 * time.Second
		baseConfig.Prefork = true
		baseConfig.WriteTimeout = 10 * time.Second
		baseConfig.IdleTimeout = 120 * time.Second

	}

	return baseConfig

}

func NewRepositories(db *pgxpool.Pool) *Repositories {
	return &Repositories{
		UserRepository:        user.NewUserRepository(db),
		ApplicationRepository: applications.NewApplicationRepository(db),
	}
}

func NewService(r *Repositories) *Services {
	return &Services{
		UserService:        user.NewUserService(r.UserRepository),
		ApplicationService: applications.NewApplicationService(r.ApplicationRepository),
	}
}
