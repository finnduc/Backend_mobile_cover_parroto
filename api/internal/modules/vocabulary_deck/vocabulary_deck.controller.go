package vocabulary_deck

import (
	"net/http"

	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/modules/vocabulary_deck/dtos/req"
	"go-cover-parroto/internal/modules/vocabulary_deck/dtos/res"
	"go-cover-parroto/internal/modules/vocabulary_deck/services"
	"go-cover-parroto/internal/utils"

	"github.com/gin-gonic/gin"
)

type VocabularyDeckController struct{ svc services.IVocabularyDeckService }

func NewVocabularyDeckController(svc services.IVocabularyDeckService) *VocabularyDeckController {
	return &VocabularyDeckController{svc: svc}
}

// List godoc
// @Summary List vocabulary decks
// @Tags vocabulary-decks
// @Accept json
// @Produce json
// @Param category_id query int false "Filter by category"
// @Param page query int false "Page number"
// @Param limit query int false "Items per page"
// @Success 200 {object} response.BaseResponse[response.PaginatedResponse[res.VocabularyDeckRes]]
// @Router /vocabulary-decks [get]
func (ctrl *VocabularyDeckController) List(c *gin.Context) {
	var q req.ListVocabularyDeckQuery
	if err := c.ShouldBindQuery(&q); err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(response.BadRequest(err.Error())))
		return
	}
	result, appErr := ctrl.svc.List(c.Request.Context(), q)
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
// @Summary Create user vocabulary deck
// @Tags vocabulary-decks
// @Accept json
// @Produce json
// @Param body body req.CreateVocabularyDeckReq true "Deck data"
// @Success 200 {object} response.BaseResponse[res.VocabularyDeckRes]
// @Failure 400 {object} response.BaseResponse[any]
// @Failure 401 {object} response.BaseResponse[any]
// @Router /vocabulary-decks [post]
// @Security BearerAuth
func (ctrl *VocabularyDeckController) Create(c *gin.Context) {
	var body req.CreateVocabularyDeckReq
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(response.BadRequest(err.Error())))
		return
	}
	deck, appErr := ctrl.svc.Create(c.Request.Context(), body)
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
// @Summary Update user vocabulary deck
// @Tags vocabulary-decks
// @Accept json
// @Produce json
// @Param id path int true "Deck ID"
// @Param body body req.UpdateVocabularyDeckReq true "Deck data"
// @Success 200 {object} response.BaseResponse[res.VocabularyDeckRes]
// @Failure 400 401 403 404 {object} response.BaseResponse[any]
// @Router /vocabulary-decks/{id} [put]
// @Security BearerAuth
func (ctrl *VocabularyDeckController) Update(c *gin.Context) {
	var uri struct {
		ID uint `uri:"id" binding:"required"`
	}
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(response.BadRequest(err.Error())))
		return
	}
	var body req.UpdateVocabularyDeckReq
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(response.BadRequest(err.Error())))
		return
	}
	deck, appErr := ctrl.svc.Update(c.Request.Context(), uri.ID, body)
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
// @Summary Delete user vocabulary deck
// @Tags vocabulary-decks
// @Accept json
// @Produce json
// @Param id path int true "Deck ID"
// @Success 200 {object} response.BaseResponse[any]
// @Failure 400 401 403 404 {object} response.BaseResponse[any]
// @Router /vocabulary-decks/{id} [delete]
// @Security BearerAuth
func (ctrl *VocabularyDeckController) Delete(c *gin.Context) {
	var uri struct {
		ID uint `uri:"id" binding:"required"`
	}
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(response.BadRequest(err.Error())))
		return
	}
	appErr := ctrl.svc.Delete(c.Request.Context(), uri.ID)
	if appErr != nil {
		c.JSON(appErr.Code, response.Fail(appErr))
		return
	}
	c.JSON(http.StatusOK, response.Success("deck deleted"))
}
