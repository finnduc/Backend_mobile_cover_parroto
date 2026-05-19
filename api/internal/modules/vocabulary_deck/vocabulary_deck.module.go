package vocabulary_deck

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"go-cover-parroto/internal/core/enums"
	"go-cover-parroto/internal/middleware"
	"go-cover-parroto/internal/modules/vocabulary_deck/repositories"
	"go-cover-parroto/internal/modules/vocabulary_deck/services"
)

func RegisterRoutes(r *gin.RouterGroup, db *gorm.DB) {
	repo := repositories.NewVocabularyDeckRepo(db)
	svc := services.NewVocabularyDeckService(repo)
	ctrl := NewVocabularyDeckController(svc)
	adminCtrl := NewVocabularyDeckAdminController(svc)

	r.GET("/vocabulary-decks", ctrl.List)

	protected := r.Group("/vocabulary-decks", middleware.ClerkAuthMiddleware())
	{
		protected.POST("", ctrl.Create)
		protected.PUT("/:id", ctrl.Update)
		protected.DELETE("/:id", ctrl.Delete)
	}

	admin := r.Group("/admin/vocabulary-decks", middleware.ClerkAuthMiddleware(), middleware.RequireRole(enums.UserRoleAdmin))
	{
		admin.GET("", adminCtrl.List)
		admin.POST("", adminCtrl.Create)
		admin.PUT("/:id", adminCtrl.Update)
		admin.DELETE("/:id", adminCtrl.Delete)
	}
}
