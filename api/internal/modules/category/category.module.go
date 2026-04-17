package category

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"go-cover-parroto/internal/modules/category/repositories"
	"go-cover-parroto/internal/modules/category/services"
)

func RegisterRoutes(r *gin.RouterGroup, db *gorm.DB) {
	repo := repositories.NewCategoryRepo(db)
	svc := services.NewCategoryService(repo)
	ctrl := NewCategoryController(svc)

	r.GET("/categories", ctrl.List)
}
