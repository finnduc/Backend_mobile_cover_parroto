package vocabulary_category

import (
	"net/http"
	"strconv"

	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/modules/vocabulary_category/dtos/req"
	"go-cover-parroto/internal/modules/vocabulary_category/dtos/res"
	"go-cover-parroto/internal/modules/vocabulary_category/services"
	"go-cover-parroto/internal/utils"

	"github.com/gin-gonic/gin"
)

type VocabularyCategoryAdminController struct {
	svc services.IVocabularyCategoryService
}

func NewVocabularyCategoryAdminController(svc services.IVocabularyCategoryService) *VocabularyCategoryAdminController {
	return &VocabularyCategoryAdminController{svc: svc}
}

// List godoc
// @Summary List vocabulary categories (admin)
// @Tags admin-vocabulary-categories
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param limit query int false "Items per page"
// @Success 200 {object} response.BaseResponse[response.PaginatedResponse[res.VocabularyCategoryRes]]
// @Failure 401 {object} response.BaseResponse[any]
// @Router /admin/vocabulary-categories [get]
// @Security BearerAuth
func (ctrl *VocabularyCategoryAdminController) List(c *gin.Context) {
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

// GetByID godoc
// @Summary Get vocabulary category by ID (admin)
// @Tags admin-vocabulary-categories
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} response.BaseResponse[res.VocabularyCategoryRes]
// @Failure 400 404 {object} response.BaseResponse[any]
// @Failure 401 {object} response.BaseResponse[any]
// @Router /admin/vocabulary-categories/{id} [get]
// @Security BearerAuth
func (ctrl *VocabularyCategoryAdminController) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(response.BadRequest("invalid id")))
		return
	}
	category, appErr := ctrl.svc.GetByID(c.Request.Context(), uint(id))
	if appErr != nil {
		c.JSON(appErr.Code, response.Fail(appErr))
		return
	}
	var result res.VocabularyCategoryRes
	if err := utils.MapToDTO(category, &result); err != nil {
		c.JSON(http.StatusInternalServerError, response.Fail(response.Internal("failed to map category")))
		return
	}
	c.JSON(http.StatusOK, response.Success(result))
}

// Create godoc
// @Summary Create vocabulary category (admin)
// @Tags admin-vocabulary-categories
// @Accept json
// @Produce json
// @Param body body req.CreateVocabularyCategoryReq true "Category data"
// @Success 200 {object} response.BaseResponse[res.VocabularyCategoryRes]
// @Failure 400 {object} response.BaseResponse[any]
// @Failure 401 {object} response.BaseResponse[any]
// @Router /admin/vocabulary-categories [post]
// @Security BearerAuth
func (ctrl *VocabularyCategoryAdminController) Create(c *gin.Context) {
	var body req.CreateVocabularyCategoryReq
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(response.BadRequest(err.Error())))
		return
	}
	category, appErr := ctrl.svc.Create(c.Request.Context(), body)
	if appErr != nil {
		c.JSON(appErr.Code, response.Fail(appErr))
		return
	}
	var result res.VocabularyCategoryRes
	if err := utils.MapToDTO(category, &result); err != nil {
		c.JSON(http.StatusInternalServerError, response.Fail(response.Internal("failed to map category")))
		return
	}
	c.JSON(http.StatusOK, response.Success(result))
}

// Update godoc
// @Summary Update vocabulary category (admin)
// @Tags admin-vocabulary-categories
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Param body body req.UpdateVocabularyCategoryReq true "Category data"
// @Success 200 {object} response.BaseResponse[res.VocabularyCategoryRes]
// @Failure 400 {object} response.BaseResponse[any]
// @Failure 401 {object} response.BaseResponse[any]
// @Router /admin/vocabulary-categories/{id} [put]
// @Security BearerAuth
func (ctrl *VocabularyCategoryAdminController) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(response.BadRequest("invalid id")))
		return
	}
	var body req.UpdateVocabularyCategoryReq
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(response.BadRequest(err.Error())))
		return
	}
	category, appErr := ctrl.svc.Update(c.Request.Context(), uint(id), body)
	if appErr != nil {
		c.JSON(appErr.Code, response.Fail(appErr))
		return
	}
	var result res.VocabularyCategoryRes
	if err := utils.MapToDTO(category, &result); err != nil {
		c.JSON(http.StatusInternalServerError, response.Fail(response.Internal("failed to map category")))
		return
	}
	c.JSON(http.StatusOK, response.Success(result))
}

// Delete godoc
// @Summary Delete vocabulary category (admin)
// @Tags admin-vocabulary-categories
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} response.BaseResponse[any]
// @Failure 400 {object} response.BaseResponse[any]
// @Failure 401 {object} response.BaseResponse[any]
// @Router /admin/vocabulary-categories/{id} [delete]
// @Security BearerAuth
func (ctrl *VocabularyCategoryAdminController) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(response.BadRequest("invalid id")))
		return
	}
	appErr := ctrl.svc.Delete(c.Request.Context(), uint(id))
	if appErr != nil {
		c.JSON(appErr.Code, response.Fail(appErr))
		return
	}
	c.JSON(http.StatusOK, response.Success("category deleted"))
}
