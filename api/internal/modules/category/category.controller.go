package category

import (
	"net/http"

	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/modules/category/services"

	_ "go-cover-parroto/internal/modules/category/dtos/res"

	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	svc services.ICategoryService
}

func NewCategoryController(svc services.ICategoryService) *CategoryController {
	return &CategoryController{svc: svc}
}

// List godoc
// @Summary List categories
// @Description Get all categories
// @Tags categories
// @Accept json
// @Produce json
// @Success 200 {object} response.BaseResponse[res.CategoryRes]
// @Router /categories [get]
func (ctrl *CategoryController) List(c *gin.Context) {
	results, appErr := ctrl.svc.ListCategories(c.Request.Context())
	if appErr != nil {
		c.JSON(appErr.Code, response.Fail(appErr))
		return
	}
	c.JSON(http.StatusOK, response.Success(results))
}
