package handlers

import (
	"net/http"

	valid "github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/hezzl_task5/internal/logger"
	"github.com/hezzl_task5/internal/models"
	"github.com/hezzl_task5/internal/usecase"
	"go.uber.org/zap"
)

type ItemsRoutes struct {
	i usecase.Item
	l *logger.Logger
}

func newItemRoutes(handler *gin.RouterGroup, t usecase.Item, l *logger.Logger) {
	r := &ItemsRoutes{t, l}

	h := handler.Group("/items")
	{
		h.GET("/get", r.get)
		h.POST("/post", r.post)
	}
}

// response for get method
type itemsResponse struct {
	Items []models.Item `json:"items"`
}

func (r *ItemsRoutes) get(c *gin.Context) {
	translations, err := r.i.Get(c.Request.Context())
	if err != nil {
		r.l.L.Error("http  - get", zap.Error(err))
		errorResponse(c, http.StatusInternalServerError, "database problems")
		return
	}

	c.JSON(http.StatusOK, itemsResponse{translations})
}

// request for post method
type createItemRequest struct {
	Name string `json:"name" valid:",required"`
}

//validate create item request
func (r *createItemRequest) validate() error {
	_, err := valid.ValidateStruct(r)
	return err
}

func (r *ItemsRoutes) post(c *gin.Context) {
	translations, err := r.i.Get(c.Request.Context())
	if err != nil {
		r.l.L.Error("http  - get", zap.Error(err))
		errorResponse(c, http.StatusInternalServerError, "database problems")
		return
	}

	c.JSON(http.StatusOK, itemsResponse{translations})
}
