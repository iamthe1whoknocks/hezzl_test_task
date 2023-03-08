package app

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/hezzl_task5/config"
	handlers "github.com/hezzl_task5/internal/handlers/http"
	"github.com/hezzl_task5/internal/logger"
	"github.com/hezzl_task5/internal/usecase"
	"github.com/hezzl_task5/internal/usecase/repo"
	"github.com/hezzl_task5/pkg/httpserver"
	"github.com/hezzl_task5/pkg/postgres"
	"go.uber.org/zap"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	l := logger.Set(cfg.Log.Level)

	// Repository
	pg, err := postgres.New(cfg.PG.URL)
	if err != nil {
		l.L.Sugar().Fatalf("app - Run - postgres.New: %w", err)
	}
	defer pg.Pool.Close()

	err = migratePostgres(l, cfg.PG.URL)
	if err != nil {
		l.L.Fatal("app - Run - migratePostgres", zap.Error(err))
	}

	// Use case
	ItemsUseCase := usecase.New(
		repo.New(pg),
	)

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
