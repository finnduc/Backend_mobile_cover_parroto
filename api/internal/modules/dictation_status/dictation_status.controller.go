package dictation_status

import (
	"net/http"

	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/modules/dictation_status/dtos/req"
	"go-cover-parroto/internal/modules/dictation_status/dtos/res"
	"go-cover-parroto/internal/modules/dictation_status/services"
	"go-cover-parroto/internal/utils"

	"github.com/gin-gonic/gin"
)

type DictationStatusController struct {
	svc services.IDictationStatusService
}

func NewDictationStatusController(svc services.IDictationStatusService) *DictationStatusController {
	return &DictationStatusController{svc: svc}
}

// Create godoc
// @Summary Mark transcript as dictation completed
// @Tags dictation-status
// @Accept json
// @Produce json
// @Param body body req.CreateDictationStatusReq true "Request body"
// @Success 200 {object} response.BaseResponse[res.DictationStatusRes]
// @Failure 400 {object} response.BaseResponse[any]
// @Failure 401 {object} response.BaseResponse[any]
// @Router /dictation-status [post]
// @Security BearerAuth
func (ctrl *DictationStatusController) Create(c *gin.Context) {
	var body req.CreateDictationStatusReq
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(response.BadRequest(err.Error())))
		return
	}
	status, appErr := ctrl.svc.Create(c.Request.Context(), body)
	if appErr != nil {
		c.JSON(appErr.Code, response.Fail(appErr))
		return
	}
	var result res.DictationStatusRes
	if err := utils.MapToDTO(status, &result); err != nil {
		c.JSON(http.StatusInternalServerError, response.Fail(response.Internal("failed to map status")))
		return
	}
	c.JSON(http.StatusOK, response.Success(result))
}

// List godoc
// @Summary List dictation status
// @Tags dictation-status
// @Accept json
// @Produce json
// @Param lesson_id query int false "Filter by lesson ID"
// @Param page query int false "Page number"
// @Param limit query int false "Items per page"
// @Success 200 {object} response.BaseResponse[response.PaginatedResponse[res.DictationStatusRes]]
// @Failure 401 {object} response.BaseResponse[any]
// @Router /dictation-status [get]
// @Security BearerAuth
func (ctrl *DictationStatusController) List(c *gin.Context) {
	var q req.ListDictationStatusQuery
	if err := c.ShouldBindQuery(&q); err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(response.BadRequest(err.Error())))
		return
	}
	result, appErr := ctrl.svc.List(c.Request.Context(), q)
	if appErr != nil {
		c.JSON(appErr.Code, response.Fail(appErr))
		return
	}
	var statuses []res.DictationStatusRes
	if err := utils.MapToDTOs(result.Data, &statuses); err != nil {
		c.JSON(http.StatusInternalServerError, response.Fail(response.Internal("failed to map statuses")))
		return
	}
	c.JSON(http.StatusOK, response.SuccessWithMeta(statuses, result.Meta))
}
