package learning_history

import (
	"net/http"
	"strconv"

	"go-cover-parroto/internal/core/enums"
	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/modules/learning_history/dtos/req"
	"go-cover-parroto/internal/modules/learning_history/services"

	"github.com/gin-gonic/gin"
)

type LearningHistoryController struct {
	svc services.ILearningHistoryService
}

func NewLearningHistoryController(svc services.ILearningHistoryService) *LearningHistoryController {
	return &LearningHistoryController{svc: svc}
}

func (ctrl *LearningHistoryController) List(c *gin.Context) {
	ctrl.list(c, "")
}

func (ctrl *LearningHistoryController) ListFinished(c *gin.Context) {
	ctrl.list(c, "finished")
}

func (ctrl *LearningHistoryController) ListUnfinished(c *gin.Context) {
	ctrl.list(c, "unfinished")
}

func (ctrl *LearningHistoryController) list(c *gin.Context, filter string) {
	userID := c.Request.Context().Value(enums.ContextKeyUserID).(string)
	result, appErr := ctrl.svc.List(c.Request.Context(), userID, filter)
	if appErr != nil {
		c.JSON(appErr.Code, response.Fail(appErr))
		return
	}
	c.JSON(http.StatusOK, response.Success(result))
}

func (ctrl *LearningHistoryController) Create(c *gin.Context) {
	userID := c.Request.Context().Value(enums.ContextKeyUserID).(string)
	var body req.CreateLearningHistoryReq
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(response.BadRequest(err.Error())))
		return
	}
	result, appErr := ctrl.svc.Create(c.Request.Context(), userID, body)
	if appErr != nil {
		c.JSON(appErr.Code, response.Fail(appErr))
		return
	}
	c.JSON(http.StatusOK, response.Success(result))
}

func (ctrl *LearningHistoryController) Summary(c *gin.Context) {
	userID := c.Request.Context().Value(enums.ContextKeyUserID).(string)
	result, appErr := ctrl.svc.Summary(c.Request.Context(), userID)
	if appErr != nil {
		c.JSON(appErr.Code, response.Fail(appErr))
		return
	}
	c.JSON(http.StatusOK, response.Success(result))
}

func (ctrl *LearningHistoryController) LessonSummary(c *gin.Context) {
	userID := c.Request.Context().Value(enums.ContextKeyUserID).(string)
	lessonID, err := strconv.ParseUint(c.Param("lessonId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(response.BadRequest("invalid lesson ID")))
		return
	}
	result, appErr := ctrl.svc.LessonSummary(c.Request.Context(), userID, uint(lessonID))
	if appErr != nil {
		c.JSON(appErr.Code, response.Fail(appErr))
		return
	}
	c.JSON(http.StatusOK, response.Success(result))
}

func (ctrl *LearningHistoryController) GetByLesson(c *gin.Context) {
	userID := c.Request.Context().Value(enums.ContextKeyUserID).(string)
	lessonID, err := strconv.ParseUint(c.Param("lessonId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(response.BadRequest("invalid lesson ID")))
		return
	}
	result, appErr := ctrl.svc.GetByLesson(c.Request.Context(), userID, uint(lessonID))
	if appErr != nil {
		c.JSON(appErr.Code, response.Fail(appErr))
		return
	}
	c.JSON(http.StatusOK, response.Success(result))
}
