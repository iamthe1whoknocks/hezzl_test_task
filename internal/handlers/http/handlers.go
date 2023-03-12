package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/iamthe1whoknocks/hezzl_test_task/internal/logger"
	"github.com/iamthe1whoknocks/hezzl_test_task/internal/usecase"
)

func NewRouter(handler *gin.Engine, i usecase.Item, l *logger.Logger) {
	// Middlewares
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	// Routers
	h := handler.Group("/v1")
	{
		newItemRoutes(h, i, l)
	}
}
