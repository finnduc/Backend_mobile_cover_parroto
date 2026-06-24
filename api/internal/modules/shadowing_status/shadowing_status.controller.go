package shadowing_status

import (
	"context"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"go-cover-parroto/internal/core/enums"
	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/modules/shadowing_status/dtos/req"
	"go-cover-parroto/internal/modules/shadowing_status/dtos/res"
	"go-cover-parroto/internal/modules/shadowing_status/services"
	"go-cover-parroto/internal/utils"

	"github.com/gin-gonic/gin"
)

type ShadowingStatusController struct {
	svc              services.IShadowingStatusService
	transcriptionSvc services.ITranscriptionService
}

func NewShadowingStatusController(svc services.IShadowingStatusService, transcriptionSvc services.ITranscriptionService) *ShadowingStatusController {
	return &ShadowingStatusController{svc: svc, transcriptionSvc: transcriptionSvc}
}

// Create godoc
// @Summary Mark transcript as shadowing completed
// @Tags shadowing-status
// @Accept json
// @Produce json
// @Param body body req.CreateShadowingStatusReq true "Request body"
// @Success 200 {object} response.BaseResponse[res.ShadowingStatusRes]
// @Failure 400 {object} response.BaseResponse[any]
// @Failure 401 {object} response.BaseResponse[any]
// @Failure 409 {object} response.BaseResponse[any]
// @Router /shadowing-status [post]
// @Security BearerAuth
func (ctrl *ShadowingStatusController) Create(c *gin.Context) {
	var body req.CreateShadowingStatusReq
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(response.BadRequest(err.Error())))
		return
	}
	userID := c.Request.Context().Value(enums.ContextKeyUserID).(string)
	status, appErr := ctrl.svc.Create(c.Request.Context(), userID, body)
	if appErr != nil {
		c.JSON(appErr.Code, response.Fail(appErr))
		return
	}
	var result res.ShadowingStatusRes
	if err := utils.MapToDTO(status, &result); err != nil {
		c.JSON(http.StatusInternalServerError, response.Fail(response.Internal("failed to map status")))
		return
	}
	c.JSON(http.StatusOK, response.Success(result))
}

// List godoc
// @Summary List shadowing status
// @Tags shadowing-status
// @Accept json
// @Produce json
// @Param lesson_id query int false "Filter by lesson ID"
// @Param page query int false "Page number"
// @Param limit query int false "Items per page"
// @Success 200 {object} response.BaseResponse[response.PaginatedResponse[res.ShadowingStatusRes]]
// @Failure 401 {object} response.BaseResponse[any]
// @Router /shadowing-status [get]
// @Security BearerAuth
func (ctrl *ShadowingStatusController) List(c *gin.Context) {
	var q req.ListShadowingStatusQuery
	if err := c.ShouldBindQuery(&q); err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(response.BadRequest(err.Error())))
		return
	}
	result, appErr := ctrl.svc.List(c.Request.Context(), q)
	if appErr != nil {
		c.JSON(appErr.Code, response.Fail(appErr))
		return
	}
	var statuses []res.ShadowingStatusRes
	if err := utils.MapToDTOs(result.Data, &statuses); err != nil {
		c.JSON(http.StatusInternalServerError, response.Fail(response.Internal("failed to map statuses")))
		return
	}
	c.JSON(http.StatusOK, response.SuccessWithMeta(statuses, result.Meta))
}

// TranscribeShadowing godoc
// @Summary Transcribe audio for shadowing exercise
// @Tags shadowing-status
// @Accept multipart/form-data
// @Produce json
// @Param audio formData file true "Audio file"
// @Success 200 {object} response.BaseResponse[res.ShadowingTranscribeRes]
// @Failure 400 {object} response.BaseResponse[any]
// @Failure 401 {object} response.BaseResponse[any]
// @Failure 500 {object} response.BaseResponse[any]
// @Router /shadowing-status/transcribe [post]
// @Security BearerAuth
func (ctrl *ShadowingStatusController) TranscribeShadowing(c *gin.Context) {
	const maxUploadSize = 10 << 20 // 10 MB
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, maxUploadSize)

	file, header, err := c.Request.FormFile("audio")
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(response.BadRequest("audio file is required")))
		return
	}
	defer file.Close()

	// Save to temp file
	tempFile, err := os.CreateTemp("", "shadowing-audio-*"+filepath.Ext(header.Filename))
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Fail(response.Internal("failed to create temp file")))
		return
	}
	defer os.Remove(tempFile.Name())
	defer tempFile.Close()

	if _, err := io.Copy(tempFile, file); err != nil {
		c.JSON(http.StatusInternalServerError, response.Fail(response.Internal("failed to save audio file")))
		return
	}
	// Transcribe using Deepgram with timeout
	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()
	transcribedText, err := ctrl.transcriptionSvc.Transcribe(ctx, tempFile.Name())
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Fail(response.Internal("transcription failed")))
		return
	}

	result := res.ShadowingTranscribeRes{
		TranscribedText: transcribedText,
	}
	c.JSON(http.StatusOK, response.Success(result))
}

// AssessPronunciation godoc
// @Summary Assess pronunciation of shadowing audio
// @Tags shadowing-status
// @Accept multipart/form-data
// @Produce json
// @Param audio formData file true "Audio file"
// @Success 200 {object} response.BaseResponse[any]
// @Failure 400 {object} response.BaseResponse[any]
// @Failure 500 {object} response.BaseResponse[any]
// @Router /pronunciation-attempts [post]
// @Security BearerAuth
func (ctrl *ShadowingStatusController) AssessPronunciation(c *gin.Context) {
	const maxUploadSize = 10 << 20 // 10 MB
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, maxUploadSize)

	file, header, err := c.Request.FormFile("audio")
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(response.BadRequest("audio file is required")))
		return
	}
	defer file.Close()

	// Save to temp file
	tempFile, err := os.CreateTemp("", "pronunciation-audio-*"+filepath.Ext(header.Filename))
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Fail(response.Internal("failed to create temp file")))
		return
	}
	defer os.Remove(tempFile.Name())
	defer tempFile.Close()

	if _, err := io.Copy(tempFile, file); err != nil {
		c.JSON(http.StatusInternalServerError, response.Fail(response.Internal("failed to save audio file")))
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()
	transcribedText, err := ctrl.transcriptionSvc.Transcribe(ctx, tempFile.Name())
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Fail(response.Internal("transcription failed")))
		return
	}

	type WeakPhoneme struct {
		Phoneme string  `json:"phoneme"`
		Score   float64 `json:"score"`
	}
	type WordItem struct {
		Word         string         `json:"word"`
		Score        float64        `json:"score"`
		Feedback     string         `json:"feedback"`
		WeakPhonemes []WeakPhoneme  `json:"weakPhonemes"`
	}
	type DetailedScores struct {
		Accuracy     float64 `json:"accuracy"`
		Fluency      float64 `json:"fluency"`
		Completeness float64 `json:"completeness"`
		Prosody      float64 `json:"prosody"`
	}
	type MockResponse struct {
		Text         string         `json:"text"`
		OverallScore float64        `json:"overallScore"`
		Feedback     string         `json:"feedback"`
		Scores       DetailedScores `json:"scores"`
		Words        []WordItem     `json:"words"`
	}

	resObj := MockResponse{
		Text:         transcribedText,
		OverallScore: 90.0,
		Feedback:     "Excellent shadowing try!",
		Scores: DetailedScores{
			Accuracy:     90.0,
			Fluency:      90.0,
			Completeness: 90.0,
			Prosody:      90.0,
		},
		Words: []WordItem{},
	}

	c.JSON(http.StatusOK, response.Success(resObj))
}

