package postgres

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4/pgxpool"
)

// Postgres struct
type Postgres struct {
	Pool    *pgxpool.Pool
	Builder squirrel.StatementBuilderType
}

func New(url string) (*Postgres, error) {

	poolConfig, err := pgxpool.ParseConfig(url)
	if err != nil {
		return nil, fmt.Errorf("postgres - NewPostgres - pgxpool.ParseConfig: %w", err)
	}

	pg := &Postgres{}
	pg.Builder = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	pg.Pool, err = pgxpool.ConnectConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, fmt.Errorf("postgres - NewPostgres - pgxpool.ConnectConfig: %w", err)
	}

	err = pg.Pool.Ping(context.Background())
	if err != nil {
		return nil, fmt.Errorf("postgres - NewPostgres - pgxpool.Ping: %w", err)
	}

	return pg, nil
}
