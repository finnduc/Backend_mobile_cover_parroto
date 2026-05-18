package learning_history

import (
	"net/http"

	"go-cover-parroto/internal/core/response"
	lhreq "go-cover-parroto/internal/modules/learning_history/dtos/req"
	"go-cover-parroto/internal/modules/learning_history/dtos/res"
	"go-cover-parroto/internal/modules/learning_history/services"
	"go-cover-parroto/internal/utils"

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
	var body lhreq.RecordHistoryReq
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(response.BadRequest(err.Error())))
		return
	}
	history, appErr := ctrl.svc.Record(c.Request.Context(), body)
	if appErr != nil {
		c.JSON(appErr.Code, response.Fail(appErr))
		return
	}
	var result res.LearningHistoryRes
	if err := utils.MapToDTO(history, &result); err != nil {
		c.JSON(http.StatusInternalServerError, response.Fail(response.Internal("failed to map history")))
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
	var q lhreq.ListHistoryQuery
	if err := c.ShouldBindQuery(&q); err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(response.BadRequest(err.Error())))
		return
	}
	result, appErr := ctrl.svc.List(c.Request.Context(), q)
	if appErr != nil {
		c.JSON(appErr.Code, response.Fail(appErr))
		return
	}
	var histories []res.LearningHistoryRes
	if err := utils.MapToDTOs(result.Data, &histories); err != nil {
		c.JSON(http.StatusInternalServerError, response.Fail(response.Internal("failed to map history")))
		return
	}
	c.JSON(http.StatusOK, response.SuccessWithMeta(histories, result.Meta))
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
	var body lhreq.GetHistoryReq
	if err := c.ShouldBindUri(&body); err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(response.BadRequest(err.Error())))
		return
	}
	history, appErr := ctrl.svc.GetByLesson(c.Request.Context(), body)
	if appErr != nil {
		c.JSON(appErr.Code, response.Fail(appErr))
		return
	}
	var result res.LearningHistoryRes
	if err := utils.MapToDTO(history, &result); err != nil {
		c.JSON(http.StatusInternalServerError, response.Fail(response.Internal("failed to map history")))
		return
	}
	c.JSON(http.StatusOK, response.Success(result))
}
