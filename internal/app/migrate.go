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
	"github.com/hezzl_task5/internal/logger"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const (
	_defaultAttempts = 20
	_defaultTimeout  = time.Second
)

func migratePostgres(l *logger.Logger, databaseURL string) error {
	if len(databaseURL) == 0 {
		return fmt.Errorf("migrate: environment variable not declared: PG_URL")
	}

	databaseURL += "?sslmode=disable"

	db, err := sql.Open("pgx", databaseURL)
	if err != nil {
		return fmt.Errorf("Migrate: connect to postgres: %s", err)
	}

	defer db.Close()

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("Migrate: get driver : %s", err)
	}

	m, err := migrate.NewWithDatabaseInstance("file:///migrations", "pgx", driver)

	if err != nil {
		return fmt.Errorf("Migrate: get driver : %s", err)
	}

	if err != nil {
		return fmt.Errorf("Migrate: postgres connect error: %s", err)
	}

	err = m.Down()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("Migrate: down error: %s", err)
	}

	err = m.Up()
	defer m.Close()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("Migrate: up error: %s", err)
	}

	if errors.Is(err, migrate.ErrNoChange) {
		l.L.Sugar().Infof("Migrate: no change")
		return nil
	}

	l.L.Sugar().Infof("Migrate: up success")
	return nil
}

// func init() {
// 	databaseURL, ok := os.LookupEnv("PG_URL")
// 	if !ok || len(databaseURL) == 0 {
// 		log.Fatalf("migrate: environment variable not declared: PG_URL")
// 	}

// 	databaseURL += "?sslmode=disable"

// 	var (
// 		attempts = _defaultAttempts
// 		err      error
// 		m        *migrate.Migrate
// 	)

// 	for attempts > 0 {
// 		m, err = migrate.New("file://migrations", databaseURL)
// 		if err == nil {
// 			break
// 		}

// 		log.Printf("Migrate: postgres is trying to connect, attempts left: %d", attempts)
// 		time.Sleep(_defaultTimeout)
// 		attempts--
// 	}

// 	if err != nil {
// 		log.Fatalf("Migrate: postgres connect error: %s", err)
// 	}

// 	err = m.Up()
// 	defer m.Close()
// 	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
// 		log.Fatalf("Migrate: up error: %s", err)
// 	}

// 	if errors.Is(err, migrate.ErrNoChange) {
// 		log.Printf("Migrate: no change")
// 		return
// 	}

// 	log.Printf("Migrate: up success")
// }
