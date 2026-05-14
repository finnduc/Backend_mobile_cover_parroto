package bookmark

import (
	"net/http"

	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/modules/bookmark/dtos/req"
	_ "go-cover-parroto/internal/modules/bookmark/dtos/res"
	"go-cover-parroto/internal/modules/bookmark/services"

	"github.com/gin-gonic/gin"
)

type BookmarkController struct {
	svc services.IBookmarkService
}

func NewBookmarkController(svc services.IBookmarkService) *BookmarkController {
	return &BookmarkController{svc: svc}
}

// List godoc
// @Summary List user bookmarks
// @Description Get the authenticated user's bookmarks with lesson details
// @Tags bookmarks
// @Accept json
// @Produce json
// @Param user_id query int false "Filter by user ID"
// @Param lesson_id query int false "Filter by lesson ID"
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Success 200 {object} response.BaseResponse[response.PaginatedResponse[res.BookmarkRes]]
// @Failure 401 {object} response.BaseResponse[any]
// @Router /bookmarks [get]
// @Security BearerAuth
func (ctrl *BookmarkController) List(c *gin.Context) {
	var q req.ListBookmarkQuery
	if err := c.ShouldBindQuery(&q); err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(response.BadRequest(err.Error())))
		return
	}
	result, appErr := ctrl.svc.List(c.Request.Context(), q)
	if appErr != nil {
		c.JSON(appErr.Code, response.Fail(appErr))
		return
	}
	c.JSON(http.StatusOK, response.SuccessWithMeta(result.Data, result.Meta))
}

// Add godoc
// @Summary Add bookmark
// @Description Add a lesson to user's bookmarks
// @Tags bookmarks
// @Accept json
// @Produce json
// @Param lessonId path int true "Lesson ID"
// @Success 200 {object} response.BaseResponse[any]
// @Failure 400 {object} response.BaseResponse[any]
// @Failure 401 {object} response.BaseResponse[any]
// @Router /bookmarks/{lessonId} [post]
// @Security BearerAuth
func (ctrl *BookmarkController) Add(c *gin.Context) {
	var body req.AddBookmarkReq
	if err := c.ShouldBindUri(&body); err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(response.BadRequest(err.Error())))
		return
	}
	appErr := ctrl.svc.AddBookmark(c.Request.Context(), body)
	if appErr != nil {
		c.JSON(appErr.Code, response.Fail(appErr))
		return
	}
	c.JSON(http.StatusOK, response.Success("Bookmark added"))
}

// Remove godoc
// @Summary Remove bookmark
// @Description Remove a lesson from user's bookmarks
// @Tags bookmarks
// @Accept json
// @Produce json
// @Param lessonId path int true "Lesson ID"
// @Success 200 {object} response.BaseResponse[any]
// @Failure 400 {object} response.BaseResponse[any]
// @Failure 401 {object} response.BaseResponse[any]
// @Router /bookmarks/{lessonId} [delete]
// @Security BearerAuth
func (ctrl *BookmarkController) Remove(c *gin.Context) {
	var body req.RemoveBookmarkReq
	if err := c.ShouldBindUri(&body); err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(response.BadRequest(err.Error())))
		return
	}
	appErr := ctrl.svc.RemoveBookmark(c.Request.Context(), body)
	if appErr != nil {
		c.JSON(appErr.Code, response.Fail(appErr))
		return
	}
	c.JSON(http.StatusOK, response.Success("Bookmark removed"))
}
