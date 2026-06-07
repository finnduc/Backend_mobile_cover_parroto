package transcript_progress

import (
	"net/http"
	"strconv"

	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/modules/transcript_progress/dtos/req"
	"go-cover-parroto/internal/modules/transcript_progress/dtos/res"
	"go-cover-parroto/internal/modules/transcript_progress/services"
	"go-cover-parroto/internal/utils"

	"github.com/gin-gonic/gin"
)

type TranscriptProgressController struct {
	svc services.ITranscriptProgressService
}

func NewTranscriptProgressController(svc services.ITranscriptProgressService) *TranscriptProgressController {
	return &TranscriptProgressController{svc: svc}
}

func (ctrl *TranscriptProgressController) Create(c *gin.Context) {
	lessonID, err := strconv.ParseUint(c.Param("lessonId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(response.BadRequest("invalid lesson ID")))
		return
	}

	var body req.CreateTranscriptProgressReq
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(response.BadRequest(err.Error())))
		return
	}

	progress, appErr := ctrl.svc.Create(c.Request.Context(), uint(lessonID), body)
	if appErr != nil {
		c.JSON(appErr.Code, response.Fail(appErr))
		return
	}

	var result res.TranscriptProgressRes
	if err := utils.MapToDTO(progress, &result); err != nil {
		c.JSON(http.StatusInternalServerError, response.Fail(response.Internal("failed to map progress")))
		return
	}

	c.JSON(http.StatusOK, response.Success(result))
}

func (ctrl *TranscriptProgressController) List(c *gin.Context) {
	lessonID, err := strconv.ParseUint(c.Param("lessonId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(response.BadRequest("invalid lesson ID")))
		return
	}

	progresses, appErr := ctrl.svc.List(c.Request.Context(), uint(lessonID))
	if appErr != nil {
		c.JSON(appErr.Code, response.Fail(appErr))
		return
	}

	var result []res.TranscriptProgressRes
	if err := utils.MapToDTOs(progresses, &result); err != nil {
		c.JSON(http.StatusInternalServerError, response.Fail(response.Internal("failed to map progresses")))
		return
	}

	c.JSON(http.StatusOK, response.Success(result))
}
