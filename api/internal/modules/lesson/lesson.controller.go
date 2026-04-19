package lesson

import (
	"net/http"
	"strconv"

	"go-cover-parroto/internal/core/database"
	"go-cover-parroto/internal/core/response"
	_ "go-cover-parroto/internal/modules/lesson/dtos/res"
	"go-cover-parroto/internal/modules/lesson/services"

	"github.com/gin-gonic/gin"
)

type LessonController struct {
	svc services.ILessonService
}

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
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	query := database.NewQuery().SetPage(page).SetLimit(limit)

	if categoryID := c.Query("category_id"); categoryID != "" {
		if id, err := strconv.ParseUint(categoryID, 10, 32); err == nil {
			query.SetFilter("category_id", uint(id))
		}
	}
	if level := c.Query("level"); level != "" {
		query.SetFilter("level", level)
	}

	results, appErr := ctrl.svc.ListLessons(c.Request.Context(), query)
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
	idParam := c.Param("lessonId")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Fail("Invalid lesson ID"))
		return
	}
	lesson, appErr := ctrl.svc.GetLesson(c.Request.Context(), uint(id))
	if appErr != nil {
		c.JSON(appErr.Code, response.Fail(appErr))
		return
	}
	c.JSON(http.StatusOK, response.Success(lesson))
}
