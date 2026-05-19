package vocabulary_category

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"go-cover-parroto/internal/core/enums"
	"go-cover-parroto/internal/middleware"
	"go-cover-parroto/internal/modules/vocabulary_category/repositories"
	"go-cover-parroto/internal/modules/vocabulary_category/services"
)

func RegisterRoutes(r *gin.RouterGroup, db *gorm.DB) {
	repo := repositories.NewVocabularyCategoryRepo(db)
	svc := services.NewVocabularyCategoryService(repo)
	ctrl := NewVocabularyCategoryController(svc)
	adminCtrl := NewVocabularyCategoryAdminController(svc)

	r.GET("/vocabulary-categories", ctrl.List)

	admin := r.Group("/admin/vocabulary-categories", middleware.ClerkAuthMiddleware(), middleware.RequireRole(enums.UserRoleAdmin))
	{
		admin.GET("", adminCtrl.List)
		admin.GET("/:id", adminCtrl.GetByID)
		admin.POST("", adminCtrl.Create)
		admin.PUT("/:id", adminCtrl.Update)
		admin.DELETE("/:id", adminCtrl.Delete)
	}
}
