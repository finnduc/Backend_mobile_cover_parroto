package bookmark

import (
	"net/http"
	"strconv"

	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/modules/bookmark/services"

	_ "go-cover-parroto/internal/modules/bookmark/dtos/res"

	"github.com/gin-gonic/gin"
)

type BookmarkController struct {
	svc services.IBookmarkService
}

func NewBookmarkController(svc services.IBookmarkService) *BookmarkController {
	return &BookmarkController{svc: svc}
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
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, response.Fail(response.Unauthorized()))
		return
	}
	lessonIdParam := c.Param("lessonId")
	lessonId, err := strconv.ParseUint(lessonIdParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Fail("Invalid lesson ID"))
		return
	}

	appErr := ctrl.svc.AddBookmark(c.Request.Context(), userID.(uint), uint(lessonId))
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
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, response.Fail(response.Unauthorized()))
		return
	}
	lessonIdParam := c.Param("lessonId")
	lessonId, err := strconv.ParseUint(lessonIdParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Fail("Invalid lesson ID"))
		return
	}

	appErr := ctrl.svc.RemoveBookmark(c.Request.Context(), userID.(uint), uint(lessonId))
	if appErr != nil {
		c.JSON(appErr.Code, response.Fail(appErr))
		return
	}
	c.JSON(http.StatusOK, response.Success("Bookmark removed"))
}
