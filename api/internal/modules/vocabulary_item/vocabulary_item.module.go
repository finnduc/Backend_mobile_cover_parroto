package vocabulary_item

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"go-cover-parroto/internal/core/enums"
	"go-cover-parroto/internal/middleware"
	deckrepos "go-cover-parroto/internal/modules/vocabulary_deck/repositories"
	"go-cover-parroto/internal/modules/vocabulary_item/repositories"
	"go-cover-parroto/internal/modules/vocabulary_item/services"
)

func RegisterRoutes(r *gin.RouterGroup, db *gorm.DB) {
	itemRepo := repositories.NewVocabularyItemRepo(db)
	deckRepo := deckrepos.NewVocabularyDeckRepo(db)
	svc := services.NewVocabularyItemService(itemRepo, deckRepo)
	ctrl := NewVocabularyItemController(svc)
	adminCtrl := NewVocabularyItemAdminController(svc)

	r.GET("/vocabulary-decks/:id/items", ctrl.List)

	protected := r.Group("", middleware.ClerkAuthMiddleware())
	{
		protected.POST("/vocabulary-decks/:id/items", ctrl.Create)

		// these should be item id, not deckId
		protected.PUT("/vocabulary-items/:id", ctrl.Update)
		protected.DELETE("/vocabulary-items/:id", ctrl.Delete)
	}

	admin := r.Group("/admin", middleware.ClerkAuthMiddleware(), middleware.RequireRole(enums.UserRoleAdmin))
	{
		admin.GET("/vocabulary-decks/:id/items", adminCtrl.List)
		admin.POST("/vocabulary-decks/:id/items", adminCtrl.Create)
		admin.PUT("/vocabulary-items/:id", adminCtrl.Update)
		admin.DELETE("/vocabulary-items/:id", adminCtrl.Delete)
	}
}
