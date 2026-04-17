package learning_history

import (
	"net/http"
	"strconv"

	"go-cover-parroto/internal/core/database"
	"go-cover-parroto/internal/core/response"
	lhreq "go-cover-parroto/internal/modules/learning_history/dtos/req"
	_ "go-cover-parroto/internal/modules/learning_history/dtos/res"
	"go-cover-parroto/internal/modules/learning_history/services"

	"github.com/gin-gonic/gin"
)

type LearningHistoryController struct {
	svc services.ILearningHistoryService
}

func NewLearningHistoryController(svc services.ILearningHistoryService) *LearningHistoryController {
	return &LearningHistoryController{svc: svc}
}

// Record godoc
// @Summary Record learning progress
// @Description Record or update a user's learning progress for a lesson
// @Tags learning-history
// @Accept json
// @Produce json
// @Param request body lhreq.RecordHistoryReq true "Record request"
// @Success 200 {object} response.BaseResponse[res.LearningHistoryRes]
// @Failure 400 {object} response.BaseResponse[any]
// @Failure 401 {object} response.BaseResponse[any]
// @Router /learning-history [post]
// @Security BearerAuth
func (ctrl *LearningHistoryController) Record(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, response.Fail(response.Unauthorized()))
		return
	}

	var body lhreq.RecordHistoryReq
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(response.BadRequest(err.Error())))
		return
	}

	result, appErr := ctrl.svc.Record(c.Request.Context(), userID.(uint), body)
	if appErr != nil {
		c.JSON(appErr.Code, response.Fail(appErr))
		return
	}
	c.JSON(http.StatusOK, response.Success(result))
}

// List godoc
// @Summary List learning history
// @Description Get the authenticated user's learning history
// @Tags learning-history
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param limit query int false "Items per page"
// @Success 200 {object} response.BaseResponse[response.PaginatedResponse[res.LearningHistoryRes]]
// @Failure 401 {object} response.BaseResponse[any]
// @Router /learning-history [get]
// @Security BearerAuth
func (ctrl *LearningHistoryController) List(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, response.Fail(response.Unauthorized()))
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	query := database.NewQuery().SetPage(page).SetLimit(limit)

	result, appErr := ctrl.svc.ListByUser(c.Request.Context(), userID.(uint), query)
	if appErr != nil {
		c.JSON(appErr.Code, response.Fail(appErr))
		return
	}
	c.JSON(http.StatusOK, response.Success(result))
}

// GetByLesson godoc
// @Summary Get history for a lesson
// @Description Get the user's learning history for a specific lesson
// @Tags learning-history
// @Accept json
// @Produce json
// @Param lessonId path int true "Lesson ID"
// @Success 200 {object} response.BaseResponse[res.LearningHistoryRes]
// @Failure 400 {object} response.BaseResponse[any]
// @Failure 401 {object} response.BaseResponse[any]
// @Failure 404 {object} response.BaseResponse[any]
// @Router /learning-history/{lessonId} [get]
// @Security BearerAuth
func (ctrl *LearningHistoryController) GetByLesson(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, response.Fail(response.Unauthorized()))
		return
	}

	lessonID, err := strconv.ParseUint(c.Param("lessonId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(response.BadRequest("invalid lesson ID")))
		return
	}

	result, appErr := ctrl.svc.GetByLesson(c.Request.Context(), userID.(uint), uint(lessonID))
	if appErr != nil {
		c.JSON(appErr.Code, response.Fail(appErr))
		return
	}
	c.JSON(http.StatusOK, response.Success(result))
}
