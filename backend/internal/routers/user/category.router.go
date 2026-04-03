package user

import (
	"go-cover-parroto/internal/controller"
	"go-cover-parroto/internal/middlewares"

	"github.com/gin-gonic/gin"
)

type CategoryRouter struct {
}

func (pr *CategoryRouter) InitCategoryRouter(r *gin.RouterGroup, categoryCtrl *controller.CategoryController) {
	category := r.Group("/categories")
	category.Use(middlewares.AuthMiddleware())
	{
		category.GET("/", categoryCtrl.List)
		category.GET("/:id/lessons", categoryCtrl.GetLessons)
	}
}
