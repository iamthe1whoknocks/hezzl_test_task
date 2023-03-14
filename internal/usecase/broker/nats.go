package broker

import (
	"context"
	"encoding/json"

	"github.com/iamthe1whoknocks/hezzl_test_task/config"
	"github.com/iamthe1whoknocks/hezzl_test_task/internal/models"
	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
)

type Broker struct {
	Conn       *nats.Conn
	Config     *config.Nats
	Logger     *zap.Logger
	ItemsBatch []*models.Item
}

func New(conn *nats.Conn, cfg *config.Nats, logger *zap.Logger) *Broker {
	return &Broker{
		Conn:       conn,
		Config:     cfg,
		Logger:     logger,
		ItemsBatch: make([]*models.Item, 0),
	}
}

func (b *Broker) Publish(ctx context.Context, subject string, data []byte) error {
	return b.Conn.Publish(subject, data)
}

func (b *Broker) GetSubject() string {
	return b.Config.Topic
}

func (b *Broker) Subscriber() {
	b.Conn.Subscribe(b.GetSubject(), func(m *nats.Msg) {
		b.Logger.Sugar().Debugf("Received a message: %s\n", string(m.Data))
		item := models.Item{}
		err := json.Unmarshal(m.Data, &item)
		if err != nil {
			b.Logger.Error("app - Subscriber - Unmarshal", zap.String("msg", string(m.Data)), zap.Error(err))
			return
		}
		b.ItemsBatch = append(b.ItemsBatch, &item)
		if len(b.ItemsBatch) == b.Config.BatchCount {

		}

	})
}
