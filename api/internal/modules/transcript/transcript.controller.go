package transcript

import (
	"net/http"
	"strconv"

	"go-cover-parroto/internal/core/response"
	_ "go-cover-parroto/internal/modules/transcript/dtos/res"
	"go-cover-parroto/internal/modules/transcript/services"

	"github.com/gin-gonic/gin"
)

type TranscriptController struct {
	svc services.ITranscriptService
}

func NewTranscriptController(svc services.ITranscriptService) *TranscriptController {
	return &TranscriptController{svc: svc}
}

// GetByLesson godoc
// @Summary Get lesson transcripts
// @Description Get all transcripts for a lesson sorted by sequence
// @Tags transcripts
// @Accept json
// @Produce json
// @Param lessonId path int true "Lesson ID"
// @Success 200 {object} response.BaseResponse[[]res.TranscriptRes]
// @Failure 400 {object} response.BaseResponse[any]
// @Failure 401 {object} response.BaseResponse[any]
// @Router /lessons/{lessonId}/transcripts [get]
// @Security BearerAuth
func (ctrl *TranscriptController) GetByLesson(c *gin.Context) {
	lessonID, err := strconv.ParseUint(c.Param("lessonId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(response.BadRequest("invalid lesson ID")))
		return
	}

	result, appErr := ctrl.svc.GetByLesson(c.Request.Context(), uint(lessonID))
	if appErr != nil {
		c.JSON(appErr.Code, response.Fail(appErr))
		return
	}
	c.JSON(http.StatusOK, response.Success(result))
}
