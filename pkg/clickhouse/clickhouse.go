// for clickhouse connection
package clickhouse

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/iamthe1whoknocks/hezzl_test_task/config"
)

type Clickhouse struct {
	DB *sql.DB
}

func New(ch *config.ClickHouse) (*Clickhouse, error) {
	// sleep for waiting clickhouse entrypoint.sh script did its job to create user and db 'hezzl'
	time.Sleep(5 * time.Second) //nolint:gomnd // test value

	db := clickhouse.OpenDB(&clickhouse.Options{
		Addr: []string{fmt.Sprintf("%s:%s", ch.Host, ch.Port)},
		Auth: clickhouse.Auth{
			Database: ch.DBName,
			Username: ch.Username,
			Password: ch.Password,
		},
	})
	err := db.Ping()
	if err != nil {
		return nil, fmt.Errorf("clickhouse - NewClickHouse - db.Ping: %w", err)
	}

	return &Clickhouse{DB: db}, nil
}
