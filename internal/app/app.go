package app

import (
	"encoding/json"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/iamthe1whoknocks/hezzl_test_task/config"
	"github.com/iamthe1whoknocks/hezzl_test_task/internal/cache"
	handlers "github.com/iamthe1whoknocks/hezzl_test_task/internal/handlers/http"
	"github.com/iamthe1whoknocks/hezzl_test_task/internal/logger"
	"github.com/iamthe1whoknocks/hezzl_test_task/internal/models"
	"github.com/iamthe1whoknocks/hezzl_test_task/internal/usecase"
	"github.com/iamthe1whoknocks/hezzl_test_task/internal/usecase/broker"
	"github.com/iamthe1whoknocks/hezzl_test_task/internal/usecase/repo"
	"github.com/iamthe1whoknocks/hezzl_test_task/pkg/clickhouse"
	"github.com/iamthe1whoknocks/hezzl_test_task/pkg/httpserver"
	"github.com/iamthe1whoknocks/hezzl_test_task/pkg/nats"
	"github.com/iamthe1whoknocks/hezzl_test_task/pkg/postgres"
	"github.com/iamthe1whoknocks/hezzl_test_task/pkg/redis"
	natsgo "github.com/nats-io/nats.go"
	"go.uber.org/zap"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	l := logger.Set(cfg.Log.Level)

	// nats
	nats, err := nats.New(&cfg.Nats)
	if err != nil {
		l.L.Sugar().Fatalf("app - Run - nats.New: %w", err)
	}
	defer nats.Conn.Close()

	// redis
	redis, err := redis.New(&cfg.Redis)
	if err != nil {
		l.L.Sugar().Fatalf("app - Run - redis.New: %w", err)
	}
	defer redis.Client.Close()

	l.L.Info("redis db started")

	// clickhouse
	ch, err := clickhouse.New(&cfg.ClickHouse)
	if err != nil {
		l.L.Sugar().Fatalf("app - Run - clickhouse.New: %w", err)
	}

	l.L.Info("clickhouse db started")

	err = migrateClickhouse(ch.DB, l)
	if err != nil {
		l.L.Fatal("app - Run - migrateClickHouse", zap.Error(err))
	}

	// postgres
	pg, err := postgres.New(cfg.PG.URL)
	if err != nil {
		l.L.Sugar().Fatalf("app - Run - postgres.New: %w", err)
	}
	defer pg.Pool.Close()

	l.L.Info("postgres db started")

	err = migratePostgres(cfg.PG.URL, l)
	if err != nil {
		l.L.Fatal("app - Run - migratePostgres", zap.Error(err))
	}

	broker := broker.New(nats.Conn, &cfg.Nats, l.L, ch.DB)

	// Use case
	ItemsUseCase := usecase.New(
		repo.New(pg, l.L),
		cache.New(redis.Client, &cfg.Redis),
		broker,
	)

	//todo: for tests only
	go func() {
		nats.Conn.Subscribe(cfg.Nats.Topic, func(m *natsgo.Msg) {
			l.L.Sugar().Debugf("Received a message: %s\n", string(m.Data))
			item := models.Item{}
			err := json.Unmarshal(m.Data, &item)
			if err != nil {
				l.L.Error("app - Subscriber - Unmarshal", zap.String("msg", string(m.Data)), zap.Error(err))
			}
		})
	}()

	// HTTP Server
	handler := gin.New()
	handlers.NewRouter(handler, ItemsUseCase, l)
	httpServer := httpserver.New(handler, cfg.HTTP.Port)

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.L.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.L.Error("app - Run - httpServer.Notify", zap.Error(err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.L.Error("app - Run - httpServer.Shutdown", zap.Error(err))
	}
}
