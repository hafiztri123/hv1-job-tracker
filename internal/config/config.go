package config

import (
	"fmt"
	"hafiztri123/hv1-job-tracker/internal/applications"
	"hafiztri123/hv1-job-tracker/internal/middleware"
	"hafiztri123/hv1-job-tracker/internal/user"
	"hafiztri123/hv1-job-tracker/internal/utils"
	"log/slog"
	"runtime/debug"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
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

func NewRouterConfig(isDev bool) fiber.Config {
	baseConfig := fiber.Config{
		AppName:      "Job Tracker v1.0",
		ServerHeader: "Fiber",
		ErrorHandler: middleware.ErrorHandler(isDev),
		BodyLimit:    10 * 10 * 1024,
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
		baseConfig.Prefork = false //change when real production
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

func NewRecoverConfig(isDev bool) recover.Config {
	if isDev {
		return recover.Config{
			EnableStackTrace: true,
			StackTraceHandler: func(c *fiber.Ctx, e interface{}) {
				slog.Error("Panic occured",
					"error", e,
					"path", c.Path(),
					"method", c.Method())

				fmt.Printf("\n📍 Stack Trace:\n%s\n\n", debug.Stack())
			},
		}
	}

	return recover.Config{
		EnableStackTrace: false,
		StackTraceHandler: func(c *fiber.Ctx, e interface{}) {
			slog.Error("panic occurred", "error", e, "path", c.Path(), "method", c.Path(), "ip", c.IP())
		},
	}

}
