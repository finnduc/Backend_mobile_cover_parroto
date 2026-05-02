package category

import (
	"net/http"
	"strconv"

	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/modules/category/dtos/req"
	_ "go-cover-parroto/internal/modules/category/dtos/res"
	"go-cover-parroto/internal/modules/category/services"
	"github.com/gin-gonic/gin"
)

type CategoryAdminController struct {
	svc services.ICategoryService
}

func NewCategoryAdminController(svc services.ICategoryService) *CategoryAdminController {
	return &CategoryAdminController{svc: svc}
}

// Create godoc
// @Summary Create category
// @Description Create a new category (admin only)
// @Tags admin-categories
// @Accept json
// @Produce json
// @Param body body req.CreateCategoryReq true "Category data"
// @Success 200 {object} response.BaseResponse[res.CategoryRes]
// @Failure 400 {object} response.BaseResponse[any]
// @Failure 401 {object} response.BaseResponse[any]
// @Router /admin/categories [post]
// @Security BearerAuth
func (ctrl *CategoryAdminController) Create(c *gin.Context) {
	var body req.CreateCategoryReq
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(response.BadRequest(err.Error())))
		return
	}
	result, appErr := ctrl.svc.Create(c.Request.Context(), body)
	if appErr != nil {
		c.JSON(appErr.Code, response.Fail(appErr))
		return
	}
	c.JSON(http.StatusOK, response.Success(result))
}

// Update godoc
// @Summary Update category
// @Description Update a category by ID (admin only)
// @Tags admin-categories
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Param body body req.UpdateCategoryReq true "Category data"
// @Success 200 {object} response.BaseResponse[res.CategoryRes]
// @Failure 400 {object} response.BaseResponse[any]
// @Failure 401 {object} response.BaseResponse[any]
// @Router /admin/categories/{id} [put]
// @Security BearerAuth
func (ctrl *CategoryAdminController) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(response.BadRequest("invalid id")))
		return
	}
	var body req.UpdateCategoryReq
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(response.BadRequest(err.Error())))
		return
	}
	result, appErr := ctrl.svc.Update(c.Request.Context(), uint(id), body)
	if appErr != nil {
		c.JSON(appErr.Code, response.Fail(appErr))
		return
	}
	c.JSON(http.StatusOK, response.Success(result))
}

// Delete godoc
// @Summary Delete category
// @Description Delete a category by ID (admin only)
// @Tags admin-categories
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} response.BaseResponse[any]
// @Failure 400 {object} response.BaseResponse[any]
// @Failure 401 {object} response.BaseResponse[any]
// @Router /admin/categories/{id} [delete]
// @Security BearerAuth
func (ctrl *CategoryAdminController) Delete(c *gin.Context) {
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
