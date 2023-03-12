package handlers

import (
	"net/http"

	valid "github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/iamthe1whoknocks/hezzl_test_task/internal/logger"
	"github.com/iamthe1whoknocks/hezzl_test_task/internal/usecase"
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
		//h.POST("/create:campaignId", r.post)
	}
}

// get handler
func (r *ItemsRoutes) get(c *gin.Context) {
	items, err := r.i.Get(c.Request.Context())
	if err != nil {
		r.l.L.Error("http  - get", zap.Error(err))
		errorResponse(c, http.StatusInternalServerError, "database problems")
		return
	}

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

//post handler
// func (r *ItemsRoutes) post(c *gin.Context) {
// 	var req createItemRequest
// 	err := c.ShouldBindJSON(&req)
// 	if err != nil {
// 		r.l.L.Error("http  - post - c.ShouldBindJSON", zap.Error(err))
// 		errorResponse(c, http.StatusBadRequest, "bad request")
// 		return
// 	}

// 	err = req.validate()
// 	if err != nil {
// 		r.l.L.Error("http  - post - dto.validate()", zap.Error(err))
// 		errorResponse(c, http.StatusBadRequest, "invalid request")
// 		return
// 	}

// 	campaignIDstr := c.Param("campaignId")
// 	campaignID, err := strconv.Atoi(campaignIDstr)
// 	if err != nil {
// 		r.l.L.Error("http  - post - strconvAtoi", zap.Error(err))
// 		errorResponse(c, http.StatusBadRequest, "invalid request")
// 		return
// 	}

// 	newItem := models.Item{
// 		CampainID: campaignID,
// 		Name:      req.Name,
// 	}

// 	item, err := r.i.Save(c.Request.Context(), newItem)

// 	c.JSON(http.StatusOK, item)
// }
