package broker

import (
	"database/sql"
	"encoding/json"
	"sync"

	"github.com/iamthe1whoknocks/hezzl_test_task/config"
	"github.com/iamthe1whoknocks/hezzl_test_task/internal/models"
	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
)

// broker struct.
type Broker struct {
	Conn       *nats.Conn
	Config     *config.Nats
	Logger     *zap.Logger
	Mu         *sync.Mutex
	ItemsBatch []*models.Item
	DB         *sql.DB
}

// new broker constructor.
func New(conn *nats.Conn, cfg *config.Nats, logger *zap.Logger, db *sql.DB) *Broker {
	return &Broker{
		Conn:       conn,
		Config:     cfg,
		Logger:     logger,
		Mu:         new(sync.Mutex),
		ItemsBatch: make([]*models.Item, 0, cfg.BatchCount),
		DB:         db,
	}
}

// Publish msg to specified subject.
func (b *Broker) Publish(subject string, data []byte) error {
	return b.Conn.Publish(subject, data)
}

// get subject.
func (b *Broker) GetSubject() string {
	return b.Config.Topic
}

// Subscribe topic and send batch of logs to clickhouse.
func (b *Broker) Subscriber() error { //nolint:gocognit /// test
	_, err := b.Conn.Subscribe(b.GetSubject(), func(m *nats.Msg) {
		b.Logger.Sugar().Debugf("app - broker - Subscriber - Received a message: %s\n", string(m.Data))
		item := models.Item{} //nolint:exhaustruct // struct to unmarshal
		err := json.Unmarshal(m.Data, &item)
		if err != nil {
			b.Logger.Error("app - Subscriber - Unmarshal", zap.String("msg", string(m.Data)), zap.Error(err))
			return
		}
		b.Mu.Lock()
		defer b.Mu.Unlock()
		b.ItemsBatch = append(b.ItemsBatch, &item)
		b.Logger.Debug("app - broker - subscribe", zap.Int("batch len", len(b.ItemsBatch)), zap.String("msg", item.Name))

		if len(b.ItemsBatch) == b.Config.BatchCount {
			var tx *sql.Tx
			b.Logger.Debug("app - broker - subscribe - sending", zap.Int("batch len", len(b.ItemsBatch)), zap.String("msg", item.Name)) //nolint:lll /// log
			tx, err = b.DB.Begin()
			if err != nil {
				b.Logger.Error("app - Subscriber - Begin", zap.String("msg", string(m.Data)), zap.Error(err))
				return
			}
			var batch *sql.Stmt

			batch, err = tx.Prepare("insert into items")
			if err != nil {
				b.Logger.Error("app - Subscriber - Prepare", zap.String("msg", string(m.Data)), zap.Error(err))
				return
			}
			defer batch.Close()
			for _, i := range b.ItemsBatch {
				_, err = batch.Exec(
					uint32(i.ID),
					uint32(i.CampainID),
					i.Name,
					i.Description,
					uint32(i.Priority),
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

			// clear batch.
			b.ItemsBatch = nil
		}
		m.Reply = "msg accepted"
		if err = m.Ack(); err != nil {
			b.Logger.Error("app - Subscriber - Ack", zap.Error(err))
		}
	})
	if err != nil {
		return err
	}
	return nil
}
