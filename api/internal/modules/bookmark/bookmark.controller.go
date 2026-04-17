package bookmark

import (
	"net/http"
	"strconv"

	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/modules/bookmark/services"

	"github.com/gin-gonic/gin"
)

type BookmarkController struct{}

func (ctrl *BookmarkController) Add(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, response.Fail("Unauthorized"))
		return
	}
	lessonIdParam := c.Param("lessonId")
	lessonId, err := strconv.ParseUint(lessonIdParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Fail("Invalid lesson ID"))
		return
	}
	err = services.AddBookmark(userID.(uint), uint(lessonId))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.Success("Bookmark added"))
}

func (ctrl *BookmarkController) Remove(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, response.Fail("Unauthorized"))
		return
	}
	lessonIdParam := c.Param("lessonId")
	lessonId, err := strconv.ParseUint(lessonIdParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Fail("Invalid lesson ID"))
		return
	}
	err = services.RemoveBookmark(userID.(uint), uint(lessonId))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.Success("Bookmark removed"))
}
