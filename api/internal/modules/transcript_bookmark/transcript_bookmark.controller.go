package transcript_bookmark

import (
	"net/http"
	"strconv"

	"go-cover-parroto/internal/core/enums"
	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/modules/transcript_bookmark/dtos/req"
	"go-cover-parroto/internal/modules/transcript_bookmark/services"

	"github.com/gin-gonic/gin"
)

type TranscriptBookmarkController struct {
	svc services.ITranscriptBookmarkService
}

func NewTranscriptBookmarkController(svc services.ITranscriptBookmarkService) *TranscriptBookmarkController {
	return &TranscriptBookmarkController{svc: svc}
}

func (ctrl *TranscriptBookmarkController) List(c *gin.Context) {
	userID := c.Request.Context().Value(enums.ContextKeyUserID).(string)
	var q req.ListTranscriptBookmarkQuery
	if err := c.ShouldBindQuery(&q); err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(response.BadRequest(err.Error())))
		return
	}
	if rawLessonID := c.Param("lessonId"); rawLessonID != "" && q.LessonID == nil {
		lessonID, err := strconv.ParseUint(rawLessonID, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, response.Fail(response.BadRequest("invalid lesson ID")))
			return
		}
		parsed := uint(lessonID)
		q.LessonID = &parsed
	}
	result, appErr := ctrl.svc.List(c.Request.Context(), userID, q)
	if appErr != nil {
		c.JSON(appErr.Code, response.Fail(appErr))
		return
	}
	c.JSON(http.StatusOK, response.Success(result))
}

func (ctrl *TranscriptBookmarkController) Create(c *gin.Context) {
	userID := c.Request.Context().Value(enums.ContextKeyUserID).(string)
	var body req.CreateTranscriptBookmarkReq
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

func (ctrl *TranscriptBookmarkController) Update(c *gin.Context) {
	userID := c.Request.Context().Value(enums.ContextKeyUserID).(string)
	transcriptID, err := strconv.ParseUint(c.Param("transcriptId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(response.BadRequest("invalid transcript ID")))
		return
	}
	var body req.UpdateTranscriptBookmarkReq
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(response.BadRequest(err.Error())))
		return
	}
	result, appErr := ctrl.svc.Update(c.Request.Context(), userID, uint(transcriptID), body)
	if appErr != nil {
		c.JSON(appErr.Code, response.Fail(appErr))
		return
	}
	c.JSON(http.StatusOK, response.Success(result))
}

func (ctrl *TranscriptBookmarkController) Delete(c *gin.Context) {
	userID := c.Request.Context().Value(enums.ContextKeyUserID).(string)
	transcriptID, err := strconv.ParseUint(c.Param("transcriptId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(response.BadRequest("invalid transcript ID")))
		return
	}
	if appErr := ctrl.svc.Delete(c.Request.Context(), userID, uint(transcriptID)); appErr != nil {
		c.JSON(appErr.Code, response.Fail(appErr))
		return
	}
	c.JSON(http.StatusOK, response.Success("Transcript bookmark deleted"))
}
