package broker

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/iamthe1whoknocks/hezzl_test_task/config"
	"github.com/iamthe1whoknocks/hezzl_test_task/internal/models"
	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
)

// broker struct
type Broker struct {
	Conn       *nats.Conn
	Config     *config.Nats
	Logger     *zap.Logger
	ItemsBatch []*models.Item
	DB         *sql.DB
}

// new broker constructor
func New(conn *nats.Conn, cfg *config.Nats, logger *zap.Logger, db *sql.DB) *Broker {
	return &Broker{
		Conn:       conn,
		Config:     cfg,
		Logger:     logger,
		ItemsBatch: make([]*models.Item, 0),
	}
}

// Publish msg to specified subject
func (b *Broker) Publish(ctx context.Context, subject string, data []byte) error {
	return b.Conn.Publish(subject, data)
}

// get subject
func (b *Broker) GetSubject() string {
	return b.Config.Topic
}

// Subscribe topic and send batch of logs to clickhouse
func (b *Broker) Subscriber() {

	b.Conn.Subscribe(b.GetSubject(), func(m *nats.Msg) {
		b.Logger.Sugar().Debugf("app - broker - Subscriber - Received a message: %s\n", string(m.Data))
		item := models.Item{}
		err := json.Unmarshal(m.Data, &item)
		if err != nil {
			b.Logger.Error("app - Subscriber - Unmarshal", zap.String("msg", string(m.Data)), zap.Error(err))
			return
		}
		b.ItemsBatch = append(b.ItemsBatch, &item)

		if len(b.ItemsBatch) == b.Config.BatchCount {
			tx, err := b.DB.Begin()
			if err != nil {
				b.Logger.Error("app - Subscriber - Begin", zap.String("msg", string(m.Data)), zap.Error(err))
				return
			}

			batch, err := tx.Prepare("insert into items")
			if err != nil {
				b.Logger.Error("app - Subscriber - Prepare", zap.String("msg", string(m.Data)), zap.Error(err))
				return
			}
			for _, i := range b.ItemsBatch {
				_, err := batch.Exec(i.ID,
					i.CampainID,
					i.Name,
					i.Description,
					i.Priority,
					i.Removed,
					i.CreatedAt)

				if err != nil {
					b.Logger.Error("app - Subscriber - Exec", zap.Any("item", i), zap.Error(err))
					continue
				}
			}
			err = tx.Commit()
			if err != nil {
				b.Logger.Error("app - Subscriber - Commit", zap.Error(err))
			}
		}
	})
}
