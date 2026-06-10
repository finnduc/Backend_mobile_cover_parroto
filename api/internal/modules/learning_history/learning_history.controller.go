package learning_history

import (
	"net/http"
	"strconv"

	"go-cover-parroto/internal/core/enums"
	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/modules/learning_history/dtos/req"
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

func (ctrl *LearningHistoryController) List(c *gin.Context) {
	userID := c.Request.Context().Value(enums.ContextKeyUserID).(string)
	items, appErr := ctrl.svc.List(c.Request.Context(), userID)
	if appErr != nil {
		c.JSON(appErr.Code, response.Fail(appErr))
		return
	}
	var result []res.LearningHistoryRes
	if err := utils.MapToDTOs(items, &result); err != nil {
		c.JSON(http.StatusInternalServerError, response.Fail(response.Internal("failed to map")))
		return
	}
	c.JSON(http.StatusOK, response.Success(result))
}

func (ctrl *LearningHistoryController) CreateOrUpdate(c *gin.Context) {
	var body req.CreateLearningHistoryReq
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(response.BadRequest(err.Error())))
		return
	}
	userID := c.Request.Context().Value(enums.ContextKeyUserID).(string)
	item, appErr := ctrl.svc.CreateOrUpdate(c.Request.Context(), userID, body)
	if appErr != nil {
		c.JSON(appErr.Code, response.Fail(appErr))
		return
	}
	var result res.LearningHistoryRes
	if err := utils.MapToDTO(item, &result); err != nil {
		c.JSON(http.StatusInternalServerError, response.Fail(response.Internal("failed to map")))
		return
	}
	c.JSON(http.StatusOK, response.Success(result))
}

func (ctrl *LearningHistoryController) ListFinished(c *gin.Context) {
	userID := c.Request.Context().Value(enums.ContextKeyUserID).(string)
	items, appErr := ctrl.svc.ListFinished(c.Request.Context(), userID)
	if appErr != nil {
		c.JSON(appErr.Code, response.Fail(appErr))
		return
	}
	var result []res.LearningHistoryRes
	if err := utils.MapToDTOs(items, &result); err != nil {
		c.JSON(http.StatusInternalServerError, response.Fail(response.Internal("failed to map")))
		return
	}
	c.JSON(http.StatusOK, response.Success(result))
}

func (ctrl *LearningHistoryController) ListUnfinished(c *gin.Context) {
	userID := c.Request.Context().Value(enums.ContextKeyUserID).(string)
	items, appErr := ctrl.svc.ListUnfinished(c.Request.Context(), userID)
	if appErr != nil {
		c.JSON(appErr.Code, response.Fail(appErr))
		return
	}
	var result []res.LearningHistoryRes
	if err := utils.MapToDTOs(items, &result); err != nil {
		c.JSON(http.StatusInternalServerError, response.Fail(response.Internal("failed to map")))
		return
	}
	c.JSON(http.StatusOK, response.Success(result))
}

func (ctrl *LearningHistoryController) Summary(c *gin.Context) {
	userID := c.Request.Context().Value(enums.ContextKeyUserID).(string)
	completed, unfinished, appErr := ctrl.svc.Summary(c.Request.Context(), userID)
	if appErr != nil {
		c.JSON(appErr.Code, response.Fail(appErr))
		return
	}
	c.JSON(http.StatusOK, response.Success(res.SummaryRes{
		CompletedCount:  completed,
		UnfinishedCount: unfinished,
	}))
}

func (ctrl *LearningHistoryController) LessonSummary(c *gin.Context) {
	lessonID, err := strconv.ParseUint(c.Param("lessonId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(response.BadRequest("invalid lesson ID")))
		return
	}
	userID := c.Request.Context().Value(enums.ContextKeyUserID).(string)
	completed, uncompleted, appErr := ctrl.svc.LessonSummary(c.Request.Context(), userID, uint(lessonID))
	if appErr != nil {
		c.JSON(appErr.Code, response.Fail(appErr))
		return
	}
	c.JSON(http.StatusOK, response.Success(res.LessonSummaryRes{
		Completed:   completed,
		Uncompleted: uncompleted,
	}))
}
