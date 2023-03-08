package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/hezzl_task5/internal/logger"
	"github.com/hezzl_task5/internal/usecase"
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
