package lesson_bookmark

import (
	"net/http"
	"strconv"

	"go-cover-parroto/internal/core/enums"
	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/modules/lesson_bookmark/dtos/req"
	"go-cover-parroto/internal/modules/lesson_bookmark/dtos/res"
	"go-cover-parroto/internal/modules/lesson_bookmark/services"
	"go-cover-parroto/internal/utils"

	"github.com/gin-gonic/gin"
)

type LessonBookmarkController struct {
	svc services.ILessonBookmarkService
}

func NewLessonBookmarkController(svc services.ILessonBookmarkService) *LessonBookmarkController {
	return &LessonBookmarkController{svc: svc}
}

func (ctrl *LessonBookmarkController) List(c *gin.Context) {
	userID := c.Request.Context().Value(enums.ContextKeyUserID).(string)

	var q req.ListLessonBookmarkQuery
	if err := c.ShouldBindQuery(&q); err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(response.BadRequest(err.Error())))
		return
	}

	result, appErr := ctrl.svc.List(c.Request.Context(), userID, q)
	if appErr != nil {
		c.JSON(appErr.Code, response.Fail(appErr))
		return
	}

	var bookmarks []res.LessonBookmarkRes
	if err := utils.MapToDTOs(result.Data, &bookmarks); err != nil {
		c.JSON(http.StatusInternalServerError, response.Fail(response.Internal("failed to map bookmarks")))
		return
	}

	c.JSON(http.StatusOK, response.SuccessWithMeta(bookmarks, result.Meta))
}

func (ctrl *LessonBookmarkController) Toggle(c *gin.Context) {
	userID := c.Request.Context().Value(enums.ContextKeyUserID).(string)

	lessonID, err := strconv.ParseUint(c.Param("lessonId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(response.BadRequest("invalid lesson ID")))
		return
	}

	bookmark, appErr := ctrl.svc.Toggle(c.Request.Context(), userID, uint(lessonID))
	if appErr != nil {
		c.JSON(appErr.Code, response.Fail(appErr))
		return
	}

	if bookmark == nil {
		c.JSON(http.StatusOK, response.Success[*res.LessonBookmarkRes](nil))
		return
	}

	var result res.LessonBookmarkRes
	if err := utils.MapToDTO(bookmark, &result); err != nil {
		c.JSON(http.StatusInternalServerError, response.Fail(response.Internal("failed to map bookmark")))
		return
	}

	c.JSON(http.StatusOK, response.Success(result))
}

func (ctrl *LessonBookmarkController) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(response.BadRequest("invalid bookmark ID")))
		return
	}

	appErr := ctrl.svc.Delete(c.Request.Context(), uint(id))
	if appErr != nil {
		c.JSON(appErr.Code, response.Fail(appErr))
		return
	}

	c.JSON(http.StatusOK, response.Success("Lesson bookmark deleted"))
}
