package main

import (
	"context"
	"hafiztri123/hv1-job-tracker/internal/config"
	"hafiztri123/hv1-job-tracker/internal/database"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	startCtx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	cfg := config.NewConfig()

	if err := database.NewDatabase(cfg, startCtx); err != nil {
		slog.Error("failed to connect to database", "error", err)
		os.Exit(1)
	}

	srv := &http.Server{
		Addr:    ":3000",
		Handler: nil,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("Failed to start server", "error", err)
			os.Exit(1)
		}

	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	slog.Info("Starting graceful shutdown...")

	endCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(endCtx); err != nil {
		slog.Error("failed to gracefully shutdown", "error", err)
		os.Exit(1)
	}
}
