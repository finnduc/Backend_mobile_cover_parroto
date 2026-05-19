package vocabulary_category

import (
	"net/http"

	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/modules/vocabulary_category/dtos/req"
	"go-cover-parroto/internal/modules/vocabulary_category/dtos/res"
	"go-cover-parroto/internal/modules/vocabulary_category/services"
	"go-cover-parroto/internal/utils"

	"github.com/gin-gonic/gin"
)

type VocabularyCategoryController struct{ svc services.IVocabularyCategoryService }

func NewVocabularyCategoryController(svc services.IVocabularyCategoryService) *VocabularyCategoryController {
	return &VocabularyCategoryController{svc: svc}
}

// List godoc
// @Summary List vocabulary categories
// @Tags vocabulary-categories
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param limit query int false "Items per page"
// @Success 200 {object} response.BaseResponse[response.PaginatedResponse[res.VocabularyCategoryRes]]
// @Failure 500 {object} response.BaseResponse[any]
// @Router /vocabulary-categories [get]
func (ctrl *VocabularyCategoryController) List(c *gin.Context) {
	var q req.ListVocabularyCategoryQuery
	if err := c.ShouldBindQuery(&q); err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(response.BadRequest(err.Error())))
		return
	}
	result, appErr := ctrl.svc.List(c.Request.Context(), q)
	if appErr != nil {
		c.JSON(appErr.Code, response.Fail(appErr))
		return
	}
	var categories []res.VocabularyCategoryRes
	if err := utils.MapToDTOs(result.Data, &categories); err != nil {
		c.JSON(http.StatusInternalServerError, response.Fail(response.Internal("failed to map categories")))
		return
	}
	c.JSON(http.StatusOK, response.SuccessWithMeta(categories, result.Meta))
}
