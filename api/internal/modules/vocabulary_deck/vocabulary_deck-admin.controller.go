package vocabulary_deck

import (
	"net/http"
	"strconv"

	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/modules/vocabulary_deck/dtos/req"
	"go-cover-parroto/internal/modules/vocabulary_deck/dtos/res"
	"go-cover-parroto/internal/modules/vocabulary_deck/services"
	"go-cover-parroto/internal/utils"

	"github.com/gin-gonic/gin"
)

type VocabularyDeckAdminController struct {
	svc services.IVocabularyDeckService
}

func NewVocabularyDeckAdminController(svc services.IVocabularyDeckService) *VocabularyDeckAdminController {
	return &VocabularyDeckAdminController{svc: svc}
}

// List godoc
// @Summary List all vocabulary decks (admin)
// @Tags admin-vocabulary-decks
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param limit query int false "Items per page"
// @Success 200 {object} response.BaseResponse[response.PaginatedResponse[res.VocabularyDeckRes]]
// @Failure 401 {object} response.BaseResponse[any]
// @Router /admin/vocabulary-decks [get]
// @Security BearerAuth
func (ctrl *VocabularyDeckAdminController) List(c *gin.Context) {
	var q req.ListVocabularyDeckQuery
	if err := c.ShouldBindQuery(&q); err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(response.BadRequest(err.Error())))
		return
	}
	result, appErr := ctrl.svc.ListDefault(c.Request.Context(), q)
	if appErr != nil {
		c.JSON(appErr.Code, response.Fail(appErr))
		return
	}
	var decks []res.VocabularyDeckRes
	if err := utils.MapToDTOs(result.Data, &decks); err != nil {
		c.JSON(http.StatusInternalServerError, response.Fail(response.Internal("failed to map decks")))
		return
	}
	c.JSON(http.StatusOK, response.SuccessWithMeta(decks, result.Meta))
}

// Create godoc
// @Summary Create system deck (admin)
// @Tags admin-vocabulary-decks
// @Accept json
// @Produce json
// @Param body body req.CreateVocabularyDeckReq true "Deck data"
// @Success 200 {object} response.BaseResponse[res.VocabularyDeckRes]
// @Failure 400 401 {object} response.BaseResponse[any]
// @Router /admin/vocabulary-decks [post]
// @Security BearerAuth
func (ctrl *VocabularyDeckAdminController) Create(c *gin.Context) {
	var body req.CreateVocabularyDeckReq
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(response.BadRequest(err.Error())))
		return
	}
	deck, appErr := ctrl.svc.CreateAsSystem(c.Request.Context(), body)
	if appErr != nil {
		c.JSON(appErr.Code, response.Fail(appErr))
		return
	}
	var result res.VocabularyDeckRes
	if err := utils.MapToDTO(deck, &result); err != nil {
		c.JSON(http.StatusInternalServerError, response.Fail(response.Internal("failed to map deck")))
		return
	}
	c.JSON(http.StatusOK, response.Success(result))
}

// Update godoc
// @Summary Update system deck (admin)
// @Tags admin-vocabulary-decks
// @Accept json
// @Produce json
// @Param id path int true "Deck ID"
// @Param body body req.UpdateVocabularyDeckReq true "Deck data"
// @Success 200 {object} response.BaseResponse[res.VocabularyDeckRes]
// @Failure 400 401 {object} response.BaseResponse[any]
// @Router /admin/vocabulary-decks/{id} [put]
// @Security BearerAuth
func (ctrl *VocabularyDeckAdminController) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(response.BadRequest("invalid id")))
		return
	}
	var body req.UpdateVocabularyDeckReq
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(response.BadRequest(err.Error())))
		return
	}
	deck, appErr := ctrl.svc.UpdateAsSystem(c.Request.Context(), uint(id), body)
	if appErr != nil {
		c.JSON(appErr.Code, response.Fail(appErr))
		return
	}
	var result res.VocabularyDeckRes
	if err := utils.MapToDTO(deck, &result); err != nil {
		c.JSON(http.StatusInternalServerError, response.Fail(response.Internal("failed to map deck")))
		return
	}
	c.JSON(http.StatusOK, response.Success(result))
}

// Delete godoc
// @Summary Delete system deck (admin)
// @Tags admin-vocabulary-decks
// @Accept json
// @Produce json
// @Param id path int true "Deck ID"
// @Success 200 {object} response.BaseResponse[any]
// @Failure 400 401 {object} response.BaseResponse[any]
// @Router /admin/vocabulary-decks/{id} [delete]
// @Security BearerAuth
func (ctrl *VocabularyDeckAdminController) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(response.BadRequest("invalid id")))
		return
	}
	appErr := ctrl.svc.DeleteAsSystem(c.Request.Context(), uint(id))
	if appErr != nil {
		c.JSON(appErr.Code, response.Fail(appErr))
		return
	}
	c.JSON(http.StatusOK, response.Success("deck deleted"))
}
