package progress

import (
	"net/http"
	"strconv"

	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/modules/progress/dtos/req"
	"go-cover-parroto/internal/modules/progress/services"

	"github.com/gin-gonic/gin"
)

type ProgressController struct{}

func (ctrl *ProgressController) Update(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, response.Fail("Unauthorized"))
		return
	}

	var body req.UpdateProgressReq
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(response.BadRequest(err.Error())))
		return
	}

	progress, err := services.UpdateProgress(userID.(uint), body.LessonID, body.DurationWatched, body.Completed)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.Success(progress))
}

func (ctrl *ProgressController) Get(c *gin.Context) {
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
	progress, err := services.GetProgress(userID.(uint), uint(lessonId))
	if err != nil {
		c.JSON(http.StatusNotFound, response.Fail("Progress not found"))
		return
	}
	c.JSON(http.StatusOK, response.Success(progress))
}
