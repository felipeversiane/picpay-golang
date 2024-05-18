package db

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Conn *pgxpool.Pool

func NewConnection(ctx context.Context, connectionString string) (*pgxpool.Pool, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	cfg, err := pgxpool.ParseConfig(connectionString)
	if err != nil {
		return nil, err
	}

	Conn, err = pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return nil, err
	}

	return Conn, nil
}
