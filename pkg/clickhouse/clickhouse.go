package clickhouse

import (
	"context"
	"fmt"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/hezzl_task5/config"
)

type Clickhouse struct {
	Conn driver.Conn
}

func New(ch *config.ClickHouse) (*Clickhouse, error) {
	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{fmt.Sprintf("%s:%s", ch.Host, ch.Port)},
		Auth: clickhouse.Auth{
			Database: ch.DbName,
			Username: ch.Username,
			Password: ch.Password,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("clickhouse - NewClikcHouse - clikchouse.Open: %w", err)
	}
	err = conn.Ping(context.Background())
	if err != nil {
		return nil, fmt.Errorf("clickhouse - NewClickHouse - conn.Ping: %w", err)
	}

	return &Clickhouse{
		Conn: conn,
	}, nil

}
