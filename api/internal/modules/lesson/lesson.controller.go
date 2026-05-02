package lesson

import (
	"net/http"

	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/modules/lesson/dtos/req"
	"go-cover-parroto/internal/modules/lesson/services"
	"github.com/gin-gonic/gin"
)

type LessonController struct{ svc services.ILessonService }

func NewLessonController(svc services.ILessonService) *LessonController {
	return &LessonController{svc: svc}
}

// List godoc
// @Summary List lessons
// @Description Get all lessons with optional filters and pagination
// @Tags lessons
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Param category_id query int false "Filter by category ID"
// @Param level query string false "Filter by level (beginner/intermediate/advanced)"
// @Success 200 {object} response.BaseResponse[response.PaginatedResponse[res.LessonRes]]
// @Failure 500 {object} response.BaseResponse[any]
// @Router /lessons [get]
func (ctrl *LessonController) List(c *gin.Context) {
	var q req.ListLessonQuery
	if err := c.ShouldBindQuery(&q); err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(response.BadRequest(err.Error())))
		return
	}
	results, appErr := ctrl.svc.List(c.Request.Context(), q)
	if appErr != nil {
		c.JSON(appErr.Code, response.Fail(appErr))
		return
	}
	c.JSON(http.StatusOK, response.Success(results))
}

// Get godoc
// @Summary Get a lesson
// @Description Get a lesson by ID
// @Tags lessons
// @Accept json
// @Produce json
// @Param lessonId path int true "Lesson ID"
// @Success 200 {object} response.BaseResponse[res.LessonRes]
// @Failure 400 {object} response.BaseResponse[any]
// @Failure 404 {object} response.BaseResponse[any]
// @Router /lessons/{lessonId} [get]
func (ctrl *LessonController) Get(c *gin.Context) {
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
	c.JSON(http.StatusOK, response.Success(lesson))
}
