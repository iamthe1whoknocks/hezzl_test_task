package app

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	// migrate tools

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/hezzl_task5/config"
	"github.com/hezzl_task5/internal/logger"
	_ "github.com/jackc/pgx/v4/stdlib"

	_ "github.com/ClickHouse/clickhouse-go"
	chm "github.com/golang-migrate/migrate/v4/database/clickhouse"
)

const (
	_defaultAttempts = 20
	_defaultTimeout  = time.Second
)

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

func clickhouseConnectionString(host, port, engine string) string {
	if engine != "" {
		return fmt.Sprintf(
			"clickhouse://%v:%v?username=user&password=password&database=db&x-multi-statement=true&x-migrations-table-engine=%v&debug=false",
			host, port, engine)
	}

	return fmt.Sprintf(
		"clickhouse://%v:%v?username=user&password=password&database=db&x-multi-statement=true&debug=false",
		host, port)
}

func migrateClickhouse(ch *config.ClickHouse, l *logger.Logger) error {
	addr := clickhouseConnectionString(ch.Host, ch.Port, "")
	conn, err := sql.Open("clickhouse", addr)
	if err != nil {
		return fmt.Errorf("app - migrateClickHouse - sql.Open :%w", err)
	}
	defer conn.Close()

	d, err := chm.WithInstance(conn, &chm.Config{})
	if err != nil {
		return fmt.Errorf("app - migrateClickHouse - clickhouse.WithInstance :%w", err)
	}

	m, err := migrate.NewWithDatabaseInstance("file:///migrations/clickhouse", "db", d)

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
