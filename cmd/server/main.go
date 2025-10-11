package main

import (
	"context"
	"fmt"
	"hafiztri123/hv1-job-tracker/internal/config"
	"hafiztri123/hv1-job-tracker/internal/database"
	"hafiztri123/hv1-job-tracker/internal/handler"
	"hafiztri123/hv1-job-tracker/internal/router"
	"hafiztri123/hv1-job-tracker/internal/utils"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		slog.Warn("godotenv failed to initialized. Using default value for env", "error", err)
	}

	startCtx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	cfg := config.NewConfig()

	db, err := database.NewDatabase(cfg, startCtx)

	if err != nil {
		slog.Error("failed to connect to database", "error", err)
		os.Exit(1)
	}

	defer db.Close()

	repos := config.NewRepositories(db.Pool)
	services := config.NewService(repos)
	handler := handler.NewHandler(services)
	app := router.NewRouter(handler)

	appPort := utils.GetEnv("APP_PORT", "3000")

	go func() {
		if err := app.Listen(fmt.Sprintf(":%s", appPort)); err != nil {
			slog.Error("failed to start server", "error", err)
			os.Exit(1)
		}
	}()

	slog.Info("app running on port", "port", appPort)

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	slog.Info("Starting graceful shutdown...")

	if err := app.ShutdownWithTimeout(30 * time.Second); err != nil {
		slog.Error("failed to gracefully shutdown", "error", err)
		os.Exit(1)
	}
}
