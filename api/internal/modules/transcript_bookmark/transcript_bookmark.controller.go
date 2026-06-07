package transcript_bookmark

import (
	"net/http"
	"strconv"

	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/modules/transcript_bookmark/dtos/req"
	"go-cover-parroto/internal/modules/transcript_bookmark/dtos/res"
	"go-cover-parroto/internal/modules/transcript_bookmark/services"
	"go-cover-parroto/internal/utils"

	"github.com/gin-gonic/gin"
)

type TranscriptBookmarkController struct {
	svc services.ITranscriptBookmarkService
}

func NewTranscriptBookmarkController(svc services.ITranscriptBookmarkService) *TranscriptBookmarkController {
	return &TranscriptBookmarkController{svc: svc}
}

// List godoc
// @Summary List transcript bookmarks for a lesson
// @Tags transcript-bookmarks
// @Accept json
// @Produce json
// @Param lessonId path int true "Lesson ID"
// @Param page query int false "Page number"
// @Param limit query int false "Items per page"
// @Success 200 {object} response.BaseResponse[any]
// @Failure 401 {object} response.BaseResponse[any]
// @Router /transcript-bookmarks/{lessonId} [get]
// @Security BearerAuth
func (ctrl *TranscriptBookmarkController) List(c *gin.Context) {
	lessonID, err := strconv.ParseUint(c.Param("lessonId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(response.BadRequest("invalid lesson ID")))
		return
	}

	var q req.ListTranscriptBookmarkQuery
	if err := c.ShouldBindQuery(&q); err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(response.BadRequest(err.Error())))
		return
	}
	lid := uint(lessonID)
	q.LessonID = &lid

	result, appErr := ctrl.svc.List(c.Request.Context(), q)
	if appErr != nil {
		c.JSON(appErr.Code, response.Fail(appErr))
		return
	}

	var bookmarks []res.TranscriptBookmarkRes
	if err := utils.MapToDTOs(result.Data, &bookmarks); err != nil {
		c.JSON(http.StatusInternalServerError, response.Fail(response.Internal("failed to map bookmarks")))
		return
	}

	c.JSON(http.StatusOK, response.SuccessWithMeta(bookmarks, result.Meta))
}

// Create godoc
// @Summary Create transcript bookmark
// @Tags transcript-bookmarks
// @Accept json
// @Produce json
// @Param body body req.CreateTranscriptBookmarkReq true "Request body"
// @Success 200 {object} response.BaseResponse[res.TranscriptBookmarkRes]
// @Failure 400 {object} response.BaseResponse[any]
// @Failure 401 {object} response.BaseResponse[any]
// @Router /transcript-bookmarks [post]
// @Security BearerAuth
func (ctrl *TranscriptBookmarkController) Create(c *gin.Context) {
	var body req.CreateTranscriptBookmarkReq
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(response.BadRequest(err.Error())))
		return
	}

	bookmark, appErr := ctrl.svc.Create(c.Request.Context(), body)
	if appErr != nil {
		c.JSON(appErr.Code, response.Fail(appErr))
		return
	}

	var result res.TranscriptBookmarkRes
	if err := utils.MapToDTO(bookmark, &result); err != nil {
		c.JSON(http.StatusInternalServerError, response.Fail(response.Internal("failed to map bookmark")))
		return
	}

	c.JSON(http.StatusOK, response.Success(result))
}

// UpdateNote godoc
// @Summary Update transcript bookmark note
// @Tags transcript-bookmarks
// @Accept json
// @Produce json
// @Param id path int true "Bookmark ID"
// @Param body body req.UpdateTranscriptBookmarkNoteReq true "Note"
// @Success 200 {object} response.BaseResponse[res.TranscriptBookmarkRes]
// @Failure 400 {object} response.BaseResponse[any]
// @Failure 401 {object} response.BaseResponse[any]
// @Failure 404 {object} response.BaseResponse[any]
// @Router /transcript-bookmarks/{id} [patch]
// @Security BearerAuth
func (ctrl *TranscriptBookmarkController) UpdateNote(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(response.BadRequest("invalid bookmark ID")))
		return
	}

	var body req.UpdateTranscriptBookmarkNoteReq
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(response.BadRequest(err.Error())))
		return
	}

	bookmark, appErr := ctrl.svc.UpdateNote(c.Request.Context(), uint(id), body)
	if appErr != nil {
		c.JSON(appErr.Code, response.Fail(appErr))
		return
	}

	var result res.TranscriptBookmarkRes
	if err := utils.MapToDTO(bookmark, &result); err != nil {
		c.JSON(http.StatusInternalServerError, response.Fail(response.Internal("failed to map bookmark")))
		return
	}

	c.JSON(http.StatusOK, response.Success(result))
}

// Delete godoc
// @Summary Delete transcript bookmark
// @Tags transcript-bookmarks
// @Accept json
// @Produce json
// @Param id path int true "Bookmark ID"
// @Success 200 {object} response.BaseResponse[any]
// @Failure 401 {object} response.BaseResponse[any]
// @Failure 404 {object} response.BaseResponse[any]
// @Router /transcript-bookmarks/{id} [delete]
// @Security BearerAuth
func (ctrl *TranscriptBookmarkController) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(response.BadRequest("invalid bookmark ID")))
		return
	}

	appErr := ctrl.svc.Delete(c.Request.Context(), uint(id))
	if appErr != nil {
		c.JSON(appErr.Code, response.Fail(appErr))
		return
	}

	c.JSON(http.StatusOK, response.Success("Transcript bookmark deleted"))
}
