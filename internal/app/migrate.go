package app

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/iamthe1whoknocks/hezzl_test_task/internal/logger"
	_ "github.com/jackc/pgx/v4/stdlib"

	"github.com/golang-migrate/migrate/v4/database/clickhouse"
)

const (
	_defaultAttempts = 20
	_defaultTimeout  = time.Second
)

// migrate postgres
func migratePostgres(databaseURL string, l *logger.Logger) error {
	if len(databaseURL) == 0 {
		return fmt.Errorf("app - migratePostgres - empty url")
	}

	databaseURL += "?sslmode=disable"

	db, err := sql.Open("pgx", databaseURL)
	if err != nil {
		return fmt.Errorf("app - migratePostgres - sql.Open : %w", err)
	}

	defer db.Close()

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("app - migratePostgres - postgres.WithInstance : %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance("file:///migrations/postgres", "pgx", driver)

	if err != nil {
		return fmt.Errorf("app - migratePostgres - migrate.NewWithDatabaseInstance : %w", err)
	}

	err = m.Down()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("app - migratePostgres - migrate.Down : %w", err)
	}

	err = m.Up()
	defer m.Close()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("app - migratePostgres - migrate.Up : %w", err)
	}

	if errors.Is(err, migrate.ErrNoChange) {
		l.L.Sugar().Infof("app - migratePostgres - migrate.Up - no change")
		return nil
	}

	l.L.Sugar().Infof("app - migratePostgres - migrate.Up - SUCCESS")
	return nil
}

// migrate clickhouse
func migrateClickhouse(db *sql.DB, l *logger.Logger) error {

	d, err := clickhouse.WithInstance(db, &clickhouse.Config{})
	if err != nil {
		return fmt.Errorf("app - migrateClickHouse - clickhouse.WithInstance :%w", err)
	}

	m, err := migrate.NewWithDatabaseInstance("file:///migrations/clickhouse", "clickhouse", d)

	if err != nil {
		return fmt.Errorf("app - migrateClickHouse - migrate.NewWithDatabaseInstance :%w", err)
	}
	err = m.Down()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("app - migrateClickHouse - migrate.Down : %w", err)
	}

	err = m.Up()
	defer m.Close()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("app - migrateClickHouse - migrate.Up : %w", err)
	}

	if errors.Is(err, migrate.ErrNoChange) {
		l.L.Sugar().Infof("app - migrateClickHouse - migrate.Up - no change")
		return nil
	}

	l.L.Sugar().Infof("app - migrateClickHouse - migrate.Up - SUCCESS")
	return nil

}
