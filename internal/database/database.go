package database

import (
	"context"
	"hafiztri123/hv1-job-tracker/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Database struct {
	Pool *pgxpool.Pool
}

func NewDatabase(cfg *config.Config, ctx context.Context) (*Database, error) {

	config, err := pgxpool.ParseConfig(cfg.DbAddr)
	if err != nil {
		return nil, err
	}

	config.MaxConns = cfg.DbMaxConns

	dbPool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		defer dbPool.Close()
		return nil, err
	}

	return &Database{Pool: dbPool}, nil
}

func (d *Database) Close() {
	if d.Pool != nil {
		d.Pool.Close()
	}
}
