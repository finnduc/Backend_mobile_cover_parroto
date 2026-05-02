package transcript

import (
	"net/http"
	"strconv"

	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/modules/transcript/dtos/req"
	_ "go-cover-parroto/internal/modules/transcript/dtos/res"
	"go-cover-parroto/internal/modules/transcript/services"

	"github.com/gin-gonic/gin"
)

type TranscriptAdminController struct {
	svc services.ITranscriptService
}

func NewTranscriptAdminController(svc services.ITranscriptService) *TranscriptAdminController {
	return &TranscriptAdminController{svc: svc}
}

// GetByID godoc
// @Summary Get transcript by ID
// @Description Get a transcript entry by ID (admin only)
// @Tags admin-transcripts
// @Accept json
// @Produce json
// @Param id path int true "Transcript ID"
// @Success 200 {object} response.BaseResponse[res.TranscriptRes]
// @Failure 400 {object} response.BaseResponse[any]
// @Failure 401 {object} response.BaseResponse[any]
// @Router /admin/transcripts/{id} [get]
// @Security BearerAuth
func (ctrl *TranscriptAdminController) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(response.BadRequest("invalid id")))
		return
	}
	result, appErr := ctrl.svc.GetByID(c.Request.Context(), uint(id))
	if appErr != nil {
		c.JSON(appErr.Code, response.Fail(appErr))
		return
	}
	c.JSON(http.StatusOK, response.Success(result))
}

// Create godoc
// @Summary Create transcript
// @Description Create a new transcript entry (admin only)
// @Tags admin-transcripts
// @Accept json
// @Produce json
// @Param body body req.CreateTranscriptReq true "Transcript data"
// @Success 200 {object} response.BaseResponse[res.TranscriptRes]
// @Failure 400 {object} response.BaseResponse[any]
// @Failure 401 {object} response.BaseResponse[any]
// @Router /admin/transcripts [post]
// @Security BearerAuth
func (ctrl *TranscriptAdminController) Create(c *gin.Context) {
	var body req.CreateTranscriptReq
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(response.BadRequest(err.Error())))
		return
	}
	result, appErr := ctrl.svc.Create(c.Request.Context(), body)
	if appErr != nil {
		c.JSON(appErr.Code, response.Fail(appErr))
		return
	}
	c.JSON(http.StatusOK, response.Success(result))
}

// Update godoc
// @Summary Update transcript
// @Description Update a transcript entry by ID (admin only)
// @Tags admin-transcripts
// @Accept json
// @Produce json
// @Param id path int true "Transcript ID"
// @Param body body req.UpdateTranscriptReq true "Transcript data"
// @Success 200 {object} response.BaseResponse[res.TranscriptRes]
// @Failure 400 {object} response.BaseResponse[any]
// @Failure 401 {object} response.BaseResponse[any]
// @Router /admin/transcripts/{id} [put]
// @Security BearerAuth
func (ctrl *TranscriptAdminController) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(response.BadRequest("invalid id")))
		return
	}
	var body req.UpdateTranscriptReq
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(response.BadRequest(err.Error())))
		return
	}
	result, appErr := ctrl.svc.Update(c.Request.Context(), uint(id), body)
	if appErr != nil {
		c.JSON(appErr.Code, response.Fail(appErr))
		return
	}
	c.JSON(http.StatusOK, response.Success(result))
}

// Delete godoc
// @Summary Delete transcript
// @Description Delete a transcript entry by ID (admin only)
// @Tags admin-transcripts
// @Accept json
// @Produce json
// @Param id path int true "Transcript ID"
// @Success 200 {object} response.BaseResponse[any]
// @Failure 400 {object} response.BaseResponse[any]
// @Failure 401 {object} response.BaseResponse[any]
// @Router /admin/transcripts/{id} [delete]
// @Security BearerAuth
func (ctrl *TranscriptAdminController) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(response.BadRequest("invalid id")))
		return
	}
	appErr := ctrl.svc.Delete(c.Request.Context(), uint(id))
	if appErr != nil {
		c.JSON(appErr.Code, response.Fail(appErr))
		return
	}
	c.JSON(http.StatusOK, response.Success("transcript deleted"))
}
