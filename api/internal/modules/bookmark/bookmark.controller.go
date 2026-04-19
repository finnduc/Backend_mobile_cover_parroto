package bookmark

import (
	"net/http"
	"strconv"

	"go-cover-parroto/internal/core/database"
	"go-cover-parroto/internal/core/response"
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
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Success 200 {object} response.BaseResponse[response.PaginatedResponse[res.BookmarkRes]]
// @Failure 401 {object} response.BaseResponse[any]
// @Router /bookmarks [get]
// @Security BearerAuth
func (ctrl *BookmarkController) List(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, response.Fail(response.Unauthorized()))
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	query := database.NewQuery().SetPage(page).SetLimit(limit)

	result, appErr := ctrl.svc.ListByUser(c.Request.Context(), userID.(uint), query)
	if appErr != nil {
		c.JSON(appErr.Code, response.Fail(appErr))
		return
	}
	c.JSON(http.StatusOK, response.Success(result))
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
