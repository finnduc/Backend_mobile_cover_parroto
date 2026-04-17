package category

import (
	"net/http"

	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/modules/category/services"

	"github.com/gin-gonic/gin"
)

type CategoryController struct{}

func (ctrl *CategoryController) List(c *gin.Context) {
	results, err := services.ListCategories()
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.Success(results))
}
