package pronunciation

import (
	"io"
	"net/http"
	"strconv"

	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/modules/pronunciation/dtos/res"
	"go-cover-parroto/internal/modules/pronunciation/services"
	"go-cover-parroto/internal/utils"

	"github.com/gin-gonic/gin"
)

type PronunciationController struct {
	svc services.IPronunciationService
}

func NewPronunciationController(svc services.IPronunciationService) *PronunciationController {
	return &PronunciationController{svc: svc}
}

func (ctrl *PronunciationController) Assess(c *gin.Context) {
	file, _, err := c.Request.FormFile("audio")
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(response.BadRequest("audio file required")))
		return
	}
	defer file.Close()

	audioData, err := io.ReadAll(file)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(response.BadRequest("failed to read audio")))
		return
	}

	referenceText := c.PostForm("referenceText")
	if referenceText == "" {
		c.JSON(http.StatusBadRequest, response.Fail(response.BadRequest("referenceText required")))
		return
	}

	lessonID, err := strconv.ParseUint(c.PostForm("lessonId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(response.BadRequest("invalid lessonId")))
		return
	}

	transcriptID, err := strconv.ParseUint(c.PostForm("transcriptId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(response.BadRequest("invalid transcriptId")))
		return
	}

	pronResult, attempt, appErr := ctrl.svc.Assess(c.Request.Context(), uint(lessonID), uint(transcriptID), referenceText, audioData)
	if appErr != nil {
		c.JSON(appErr.Code, response.Fail(appErr))
		return
	}

	result := res.PronunciationRes{
		Text:         pronResult.Text,
		OverallScore: pronResult.OverallScore,
		Scores: res.PronunciationScores{
			Accuracy:     pronResult.Scores.Accuracy,
			Fluency:      pronResult.Scores.Fluency,
			Completeness: pronResult.Scores.Completeness,
			Prosody:      pronResult.Scores.Prosody,
		},
		Feedback: pronResult.Feedback,
		Attempt: res.AttemptInfo{
			ID:           attempt.ID,
			UserID:       attempt.UserID,
			LessonID:     attempt.LessonID,
			TranscriptID: attempt.TranscriptID,
			CreatedAt:    attempt.CreatedAt.Format("2006-01-02T15:04:05Z"),
		},
	}

	for _, w := range pronResult.Words {
		wordRes := res.WordResult{
			Word:     w.Word,
			Score:    w.Score,
			Feedback: w.Feedback,
		}
		for _, p := range w.WeakPhonemes {
			wordRes.WeakPhonemes = append(wordRes.WeakPhonemes, res.PhonemeResult{
				Phoneme: p.Phoneme,
				Score:   p.Score,
			})
		}
		result.Words = append(result.Words, wordRes)
	}

	c.JSON(http.StatusOK, response.Success(result))
}

func (ctrl *PronunciationController) DeleteAttempt(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("attemptId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(response.BadRequest("invalid attempt ID")))
		return
	}
	appErr := ctrl.svc.DeleteAttempt(c.Request.Context(), uint(id))
	if appErr != nil {
		c.JSON(appErr.Code, response.Fail(appErr))
		return
	}
	c.JSON(http.StatusOK, response.Success("Pronunciation attempt deleted"))
}

func (ctrl *PronunciationController) ListProgress(c *gin.Context) {
	lessonID, err := strconv.ParseUint(c.Param("lessonId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(response.BadRequest("invalid lesson ID")))
		return
	}
	items, appErr := ctrl.svc.ListProgress(c.Request.Context(), uint(lessonID))
	if appErr != nil {
		c.JSON(appErr.Code, response.Fail(appErr))
		return
	}
	var result []res.PronunciationProgressRes
	if err := utils.MapToDTOs(items, &result); err != nil {
		c.JSON(http.StatusInternalServerError, response.Fail(response.Internal("failed to map")))
		return
	}
	c.JSON(http.StatusOK, response.Success(result))
}

func (ctrl *PronunciationController) UpdateProgress(c *gin.Context) {
	transcriptID, err := strconv.ParseUint(c.Param("transcriptId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(response.BadRequest("invalid transcript ID")))
		return
	}
	item, appErr := ctrl.svc.UpdateProgress(c.Request.Context(), uint(transcriptID))
	if appErr != nil {
		c.JSON(appErr.Code, response.Fail(appErr))
		return
	}
	var result res.PronunciationProgressRes
	if err := utils.MapToDTO(item, &result); err != nil {
		c.JSON(http.StatusInternalServerError, response.Fail(response.Internal("failed to map")))
		return
	}
	c.JSON(http.StatusOK, response.Success(result))
}
