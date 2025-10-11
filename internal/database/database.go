package database

import (
	"context"
	"hafiztri123/hv1-job-tracker/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewDatabase(cfg *config.Config, ctx context.Context) error {

	config, err := pgxpool.ParseConfig(cfg.DbAddr)
	if err != nil {
		return err
	}

	config.MaxConns = cfg.DbMaxConns

	dbPool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return err
	}

	defer dbPool.Close()

	return nil
}
