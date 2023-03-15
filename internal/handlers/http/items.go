package handlers

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	valid "github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/iamthe1whoknocks/hezzl_test_task/internal/logger"
	"github.com/iamthe1whoknocks/hezzl_test_task/internal/models"
	"github.com/iamthe1whoknocks/hezzl_test_task/internal/usecase"
	"github.com/jackc/pgx/v4"
	"github.com/vmihailenco/msgpack/v5"
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
		h.GET("/list", r.get)
		h.POST("/create", r.post)
		h.DELETE("/remove", r.delete)
		h.PATCH("/update", r.update)
	}
}

// get handler
func (r *ItemsRoutes) get(c *gin.Context) {

	// find items in cache
	b, err := r.i.GetCache(c.Request.Context(), "get")
	if err == nil && b != nil {
		items := make([]models.Item, 0)
		err = msgpack.Unmarshal(b, &items)
		if err != nil {
			r.l.L.Error("http  - get - SetCache", zap.Error(err))
			errorResponse(c, http.StatusInternalServerError, "internal error")
			return
		}
		r.l.L.Debug("http - get - got from cache")
		c.JSON(http.StatusOK, items)
		return
	}

	items, err := r.i.Get(c.Request.Context())
	if err != nil {
		r.l.L.Error("http  - get", zap.Error(err))
		errorResponse(c, http.StatusInternalServerError, "database problems")
		return
	}

	if len(items) == 0 {
		c.JSON(http.StatusOK, "items list is empty")
		return
	}

	b, err = msgpack.Marshal(items)
	if err != nil {
		r.l.L.Error("http  - get - msgpack.Marshal", zap.Error(err))
		errorResponse(c, http.StatusInternalServerError, "internal error")
		return
	}

	// key is the same because our get method lists all items, not by one
	err = r.i.SetCache(c.Request.Context(), "get", b)
	if err != nil {
		r.l.L.Error("http  - get - SetCache", zap.Error(err))
		errorResponse(c, http.StatusInternalServerError, "internal error")
		return
	}
	r.l.L.Debug("http - get - got from db")
	c.JSON(http.StatusOK, items)
}

// request for post method
type createItemRequest struct {
	Name string `json:"name" valid:",required"`
}

// validate create item request
func (r *createItemRequest) validate() error {
	_, err := valid.ValidateStruct(r)
	return err
}

// post handler
func (r *ItemsRoutes) post(c *gin.Context) {
	var req createItemRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		r.l.L.Error("http  - post - c.ShouldBindJSON", zap.Error(err))
		errorResponse(c, http.StatusBadRequest, "bad request")
		return
	}

	err = req.validate()
	if err != nil {
		r.l.L.Error("http  - post - dto.validate()", zap.Error(err))
		errorResponse(c, http.StatusBadRequest, "invalid request")
		return
	}

	campaignIDstr := c.Query("campaignId")
	campaignID, err := strconv.Atoi(campaignIDstr)
	if err != nil {
		r.l.L.Error("http  - post - strconvAtoi", zap.Error(err))
		errorResponse(c, http.StatusBadRequest, "invalid request")
		return
	}

	newItem := models.Item{
		CampainID: campaignID,
		Name:      req.Name,
	}

	item, err := r.i.Save(c.Request.Context(), &newItem)
	if err != nil {
		if strings.Contains(err.Error(), "violates foreign key constraint") {
			r.l.L.Error("http  - post - r.i.Save - foreign key constraint", zap.Error(err))
			errorResponse(c, http.StatusBadRequest, "bad request")
			return
		}
		r.l.L.Error("http  - post - r.i.Save", zap.Error(err))
		errorResponse(c, http.StatusInternalServerError, "internal error")
		return
	}

	err = r.i.InvalidateCache(c.Request.Context(), "get")
	if err != nil {
		r.l.L.Error("http  - post - r.i.InvalidateCache", zap.Error(err))
		errorResponse(c, http.StatusInternalServerError, "internal error")
		return
	}

	c.JSON(http.StatusOK, item)
}

type deleteResponse struct {
	ID         int  `json:"id"`
	CampaignID int  `json:"campaign_id"`
	Removed    bool `json:"removed"`
}

// delete handler
func (r *ItemsRoutes) delete(c *gin.Context) {
	campaignIDstr := c.Query("campaignId")
	campaignID, err := strconv.Atoi(campaignIDstr)
	if err != nil {
		r.l.L.Error("http  - delete - campaignId -  strconvAtoi", zap.Error(err))
		errorResponse(c, http.StatusBadRequest, "invalid request")
		return
	}

	IDstr := c.Query("id")
	id, err := strconv.Atoi(IDstr)
	if err != nil {
		r.l.L.Error("http  - delete - id - strconvAtoi", zap.Error(err))
		errorResponse(c, http.StatusBadRequest, "invalid request")
		return
	}

	isDeleted, err := r.i.Delete(c.Request.Context(), id, campaignID)
	if err != nil {
		if errors.Unwrap(err) == pgx.ErrNoRows {
			r.l.L.Error("http  - delete - r.i.Delete - sql.ErrNoRows", zap.Error(err))
			c.JSON(http.StatusOK, "item was not found")
			return
		} else {
			r.l.L.Error("http  - delete - r.i.Delete", zap.Error(err))
			errorResponse(c, http.StatusInternalServerError, "internal error")
			return
		}
	}

	err = r.i.InvalidateCache(c.Request.Context(), "get")
	if err != nil {
		r.l.L.Error("http  - delete - r.i.InvalidateCache", zap.Error(err))
		errorResponse(c, http.StatusInternalServerError, "internal error")
		return
	}

	c.JSON(http.StatusOK, deleteResponse{
		ID:         id,
		CampaignID: campaignID,
		Removed:    isDeleted,
	})
}

type updateRequest struct {
	Name        string `json:"name" valid:",required"`
	Description string `json:"description"`
}

// validate update request
func (r *updateRequest) validate() error {
	_, err := valid.ValidateStruct(r)
	return err
}

// update handler
func (r *ItemsRoutes) update(c *gin.Context) {
	campaignIDstr := c.Query("campaignId")
	campaignID, err := strconv.Atoi(campaignIDstr)
	if err != nil {
		r.l.L.Error("http  - delete - campaignId -  strconvAtoi", zap.Error(err))
		errorResponse(c, http.StatusBadRequest, "invalid request")
		return
	}

	IDstr := c.Query("id")
	id, err := strconv.Atoi(IDstr)
	if err != nil {
		r.l.L.Error("http  - delete - id - strconvAtoi", zap.Error(err))
		errorResponse(c, http.StatusBadRequest, "invalid request")
		return
	}

	var dto updateRequest

	err = c.ShouldBindJSON(&dto)
	if err != nil {
		r.l.L.Error("http  - update - c.ShouldBindJSON", zap.Error(err))
		errorResponse(c, http.StatusBadRequest, "bad request")
		return
	}

	err = dto.validate()
	if err != nil {
		r.l.L.Error("http  - update - dto.validate()", zap.Error(err))
		errorResponse(c, http.StatusBadRequest, "bad request")
		return
	}

	item := &models.Item{
		ID:          id,
		CampainID:   campaignID,
		Name:        dto.Name,
		Description: dto.Description,
	}

	item, err = r.i.Update(c.Request.Context(), item)
	if err != nil {
		if errors.Unwrap(err) == pgx.ErrNoRows {
			r.l.L.Error("http  - update - r.i.Update - sql.ErrNoRows", zap.Error(err))
			c.JSON(http.StatusOK, "item was not found")
			return
		} else {
			r.l.L.Error("http  - update - r.i.Update", zap.Error(err))
			errorResponse(c, http.StatusInternalServerError, "internal error")
			return
		}

	}

	err = r.i.InvalidateCache(c.Request.Context(), "get")
	if err != nil {
		r.l.L.Error("http  - update - r.i.InvalidateCache", zap.Error(err))
		errorResponse(c, http.StatusInternalServerError, "internal error")
		return
	}

	c.JSON(http.StatusOK, item)
}
