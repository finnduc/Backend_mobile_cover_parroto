package lesson

import (
	"net/http"
	"strconv"

	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/modules/lesson/services"

	"github.com/gin-gonic/gin"
)

type LessonController struct{}

func (ctrl *LessonController) List(c *gin.Context) {
	results, err := services.ListLessons()
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.Success(results))
}

func (ctrl *LessonController) Get(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Fail("Invalid lesson ID"))
		return
	}
	lesson, err := services.GetLesson(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, response.Fail("Lesson not found"))
		return
	}
	c.JSON(http.StatusOK, response.Success(lesson))
}
