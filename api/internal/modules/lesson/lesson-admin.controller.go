package lesson

import (
	"net/http"
	"strconv"

	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/modules/lesson/dtos/req"
	"go-cover-parroto/internal/modules/lesson/dtos/res"
	"go-cover-parroto/internal/modules/lesson/services"
	"go-cover-parroto/internal/utils"

	"github.com/gin-gonic/gin"
)

type LessonAdminController struct {
	svc services.ILessonService
}

func NewLessonAdminController(svc services.ILessonService) *LessonAdminController {
	return &LessonAdminController{svc: svc}
}

// List godoc
// @Summary List lessons
// @Description List all lessons with pagination (admin only)
// @Tags admin-lessons
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Param category_id query int false "Filter by category ID"
// @Param level query string false "Filter by level"
// @Success 200 {object} response.BaseResponse[response.PaginatedResponse[res.LessonRes]]
// @Failure 401 {object} response.BaseResponse[any]
// @Router /admin/lessons [get]
// @Security BearerAuth
func (ctrl *LessonAdminController) List(c *gin.Context) {
	var q req.ListLessonQuery
	if err := c.ShouldBindQuery(&q); err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(response.BadRequest(err.Error())))
		return
	}
	result, appErr := ctrl.svc.List(c.Request.Context(), q)
	if appErr != nil {
		c.JSON(appErr.Code, response.Fail(appErr))
		return
	}
	var lessons []res.LessonRes
	if err := utils.MapToDTOs(result.Data, &lessons); err != nil {
		c.JSON(http.StatusInternalServerError, response.Fail(response.Internal("failed to map lessons")))
		return
	}
	c.JSON(http.StatusOK, response.SuccessWithMeta(lessons, result.Meta))
}

// Get godoc
// @Summary Get lesson by ID
// @Description Get a lesson by ID (admin only)
// @Tags admin-lessons
// @Accept json
// @Produce json
// @Param id path int true "Lesson ID"
// @Success 200 {object} response.BaseResponse[res.LessonRes]
// @Failure 400 {object} response.BaseResponse[any]
// @Failure 401 {object} response.BaseResponse[any]
// @Router /admin/lessons/{id} [get]
// @Security BearerAuth
func (ctrl *LessonAdminController) Get(c *gin.Context) {
	var body req.GetLessonReq
	if err := c.ShouldBindUri(&body); err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(response.BadRequest(err.Error())))
		return
	}
	lesson, appErr := ctrl.svc.Get(c.Request.Context(), body)
	if appErr != nil {
		c.JSON(appErr.Code, response.Fail(appErr))
		return
	}
	var result res.LessonRes
	if err := utils.MapToDTO(lesson, &result); err != nil {
		c.JSON(http.StatusInternalServerError, response.Fail(response.Internal("failed to map lesson")))
		return
	}
	c.JSON(http.StatusOK, response.Success(result))
}

// Create godoc
// @Summary Create lesson
// @Description Create a new lesson (admin only)
// @Tags admin-lessons
// @Accept json
// @Produce json
// @Param body body req.CreateLessonReq true "Lesson data"
// @Success 200 {object} response.BaseResponse[res.LessonRes]
// @Failure 400 {object} response.BaseResponse[any]
// @Failure 401 {object} response.BaseResponse[any]
// @Router /admin/lessons [post]
// @Security BearerAuth
func (ctrl *LessonAdminController) Create(c *gin.Context) {
	var body req.CreateLessonReq
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(response.BadRequest(err.Error())))
		return
	}
	lesson, appErr := ctrl.svc.Create(c.Request.Context(), body)
	if appErr != nil {
		c.JSON(appErr.Code, response.Fail(appErr))
		return
	}
	var result res.LessonRes
	if err := utils.MapToDTO(lesson, &result); err != nil {
		c.JSON(http.StatusInternalServerError, response.Fail(response.Internal("failed to map lesson")))
		return
	}
	c.JSON(http.StatusOK, response.Success(result))
}

// Update godoc
// @Summary Update lesson
// @Description Update a lesson by ID (admin only)
// @Tags admin-lessons
// @Accept json
// @Produce json
// @Param id path int true "Lesson ID"
// @Param body body req.UpdateLessonReq true "Lesson data"
// @Success 200 {object} response.BaseResponse[res.LessonRes]
// @Failure 400 {object} response.BaseResponse[any]
// @Failure 401 {object} response.BaseResponse[any]
// @Router /admin/lessons/{id} [put]
// @Security BearerAuth
func (ctrl *LessonAdminController) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("lessonId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(response.BadRequest("invalid id")))
		return
	}
	var body req.UpdateLessonReq
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(response.BadRequest(err.Error())))
		return
	}
	lesson, appErr := ctrl.svc.Update(c.Request.Context(), uint(id), body)
	if appErr != nil {
		c.JSON(appErr.Code, response.Fail(appErr))
		return
	}
	var result res.LessonRes
	if err := utils.MapToDTO(lesson, &result); err != nil {
		c.JSON(http.StatusInternalServerError, response.Fail(response.Internal("failed to map lesson")))
		return
	}
	c.JSON(http.StatusOK, response.Success(result))
}

// Delete godoc
// @Summary Delete lesson
// @Description Delete a lesson by ID (admin only)
// @Tags admin-lessons
// @Accept json
// @Produce json
// @Param id path int true "Lesson ID"
// @Success 200 {object} response.BaseResponse[any]
// @Failure 400 {object} response.BaseResponse[any]
// @Failure 401 {object} response.BaseResponse[any]
// @Router /admin/lessons/{id} [delete]
// @Security BearerAuth
func (ctrl *LessonAdminController) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("lessonId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(response.BadRequest("invalid id")))
		return
	}
	appErr := ctrl.svc.Delete(c.Request.Context(), uint(id))
	if appErr != nil {
		c.JSON(appErr.Code, response.Fail(appErr))
		return
	}
	c.JSON(http.StatusOK, response.Success("lesson deleted"))
}
