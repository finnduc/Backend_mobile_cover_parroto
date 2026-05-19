package shadowing_status

import (
	"net/http"

	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/modules/shadowing_status/dtos/req"
	"go-cover-parroto/internal/modules/shadowing_status/dtos/res"
	"go-cover-parroto/internal/modules/shadowing_status/services"
	"go-cover-parroto/internal/utils"

	"github.com/gin-gonic/gin"
)

type ShadowingStatusController struct {
	svc services.IShadowingStatusService
}

func NewShadowingStatusController(svc services.IShadowingStatusService) *ShadowingStatusController {
	return &ShadowingStatusController{svc: svc}
}

// Create godoc
// @Summary Mark transcript as shadowing completed
// @Tags shadowing-status
// @Accept json
// @Produce json
// @Param transcriptId path int true "Transcript ID"
// @Success 200 {object} response.BaseResponse[res.ShadowingStatusRes]
// @Failure 400 {object} response.BaseResponse[any]
// @Failure 401 {object} response.BaseResponse[any]
// @Failure 409 {object} response.BaseResponse[any]
// @Router /shadowing-status/{transcriptId} [post]
// @Security BearerAuth
func (ctrl *ShadowingStatusController) Create(c *gin.Context) {
	var body req.CreateShadowingStatusReq
	if err := c.ShouldBindUri(&body); err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(response.BadRequest(err.Error())))
		return
	}
	status, appErr := ctrl.svc.Create(c.Request.Context(), body)
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
