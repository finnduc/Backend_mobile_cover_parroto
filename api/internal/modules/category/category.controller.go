package category

import (
	"net/http"
	"strconv"

	"go-cover-parroto/internal/core/database"
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
// @Description Get all categories with pagination
// @Tags categories
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Success 200 {object} response.BaseResponse[response.PaginatedResponse[res.CategoryRes]]
// @Failure 500 {object} response.BaseResponse[any]
// @Router /categories [get]
func (ctrl *CategoryController) List(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	query := database.NewQuery().SetPage(page).SetLimit(limit)
	results, appErr := ctrl.svc.ListCategories(c.Request.Context(), query)
	if appErr != nil {
		c.JSON(appErr.Code, response.Fail(appErr))
		return
	}
	c.JSON(http.StatusOK, response.Success(results))
}
