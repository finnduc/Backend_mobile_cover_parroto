package vocabulary_item

import (
	"net/http"
	"strconv"

	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/modules/vocabulary_item/dtos/req"
	"go-cover-parroto/internal/modules/vocabulary_item/dtos/res"
	"go-cover-parroto/internal/modules/vocabulary_item/services"
	"go-cover-parroto/internal/utils"

	"github.com/gin-gonic/gin"
)

type VocabularyItemAdminController struct {
	svc services.IVocabularyItemService
}

func NewVocabularyItemAdminController(svc services.IVocabularyItemService) *VocabularyItemAdminController {
	return &VocabularyItemAdminController{svc: svc}
}

// List godoc
// @Summary List items in a system deck (admin)
// @Tags admin-vocabulary-items
// @Accept json
// @Produce json
// @Param deckId path int true "Deck ID"
// @Param page query int false "Page number"
// @Param limit query int false "Items per page"
// @Success 200 {object} response.BaseResponse[response.PaginatedResponse[res.VocabularyItemRes]]
// @Router /admin/vocabulary-decks/{deckId}/items [get]
// @Security BearerAuth
func (ctrl *VocabularyItemAdminController) List(c *gin.Context) {
	var uri struct {
		DeckID uint `uri:"deckId" binding:"required"`
	}
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(response.BadRequest(err.Error())))
		return
	}
	var q req.ListVocabularyItemQuery
	if err := c.ShouldBindQuery(&q); err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(response.BadRequest(err.Error())))
		return
	}
	q.DeckID = &uri.DeckID

	result, appErr := ctrl.svc.List(c.Request.Context(), q)
	if appErr != nil {
		c.JSON(appErr.Code, response.Fail(appErr))
		return
	}
	var items []res.VocabularyItemRes
	if err := utils.MapToDTOs(result.Data, &items); err != nil {
		c.JSON(http.StatusInternalServerError, response.Fail(response.Internal("failed to map items")))
		return
	}
	c.JSON(http.StatusOK, response.SuccessWithMeta(items, result.Meta))
}

// Create godoc
// @Summary Add item to system deck (admin)
// @Tags admin-vocabulary-items
// @Accept json
// @Produce json
// @Param deckId path int true "Deck ID"
// @Param body body req.CreateVocabularyItemReq true "Item data"
// @Success 200 {object} response.BaseResponse[res.VocabularyItemRes]
// @Failure 400 401 {object} response.BaseResponse[any]
// @Router /admin/vocabulary-decks/{deckId}/items [post]
// @Security BearerAuth
func (ctrl *VocabularyItemAdminController) Create(c *gin.Context) {
	var uri struct {
		DeckID uint `uri:"deckId" binding:"required"`
	}
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(response.BadRequest(err.Error())))
		return
	}
	var body req.CreateVocabularyItemReq
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(response.BadRequest(err.Error())))
		return
	}

	adminBody := req.CreateVocabularyItemFromDeckReq{
		DeckID:           uri.DeckID,
		Phrase:           body.Phrase,
		NormalizedPhrase: body.NormalizedPhrase,
		Meaning:          body.Meaning,
		ExampleSentence:  body.ExampleSentence,
		Note:             body.Note,
		LessonID:         body.LessonID,
		TranscriptID:     body.TranscriptID,
	}

	item, appErr := ctrl.svc.Create(c.Request.Context(), adminBody)
	if appErr != nil {
		c.JSON(appErr.Code, response.Fail(appErr))
		return
	}
	var result res.VocabularyItemRes
	if err := utils.MapToDTO(item, &result); err != nil {
		c.JSON(http.StatusInternalServerError, response.Fail(response.Internal("failed to map item")))
		return
	}
	c.JSON(http.StatusOK, response.Success(result))
}

// Update godoc
// @Summary Update item in system deck (admin)
// @Tags admin-vocabulary-items
// @Accept json
// @Produce json
// @Param id path int true "Item ID"
// @Param body body req.UpdateVocabularyItemReq true "Item data"
// @Success 200 {object} response.BaseResponse[res.VocabularyItemRes]
// @Failure 400 401 {object} response.BaseResponse[any]
// @Router /admin/vocabulary-items/{id} [put]
// @Security BearerAuth
func (ctrl *VocabularyItemAdminController) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(response.BadRequest("invalid id")))
		return
	}
	var body req.UpdateVocabularyItemReq
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(response.BadRequest(err.Error())))
		return
	}
	item, appErr := ctrl.svc.Update(c.Request.Context(), uint(id), body)
	if appErr != nil {
		c.JSON(appErr.Code, response.Fail(appErr))
		return
	}
	var result res.VocabularyItemRes
	if err := utils.MapToDTO(item, &result); err != nil {
		c.JSON(http.StatusInternalServerError, response.Fail(response.Internal("failed to map item")))
		return
	}
	c.JSON(http.StatusOK, response.Success(result))
}

// Delete godoc
// @Summary Delete item from system deck (admin)
// @Tags admin-vocabulary-items
// @Accept json
// @Produce json
// @Param id path int true "Item ID"
// @Success 200 {object} response.BaseResponse[any]
// @Failure 400 401 {object} response.BaseResponse[any]
// @Router /admin/vocabulary-items/{id} [delete]
// @Security BearerAuth
func (ctrl *VocabularyItemAdminController) Delete(c *gin.Context) {
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
	c.JSON(http.StatusOK, response.Success("item deleted"))
}
