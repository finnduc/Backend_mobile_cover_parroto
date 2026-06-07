package transcript_bookmark

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"go-cover-parroto/internal/middleware"
	"go-cover-parroto/internal/modules/transcript_bookmark/repositories"
	"go-cover-parroto/internal/modules/transcript_bookmark/services"
)

func RegisterRoutes(r *gin.RouterGroup, db *gorm.DB) {
	repo := repositories.NewTranscriptBookmarkRepo(db)
	svc := services.NewTranscriptBookmarkService(repo)
	ctrl := NewTranscriptBookmarkController(svc)

	protected := r.Group("/transcript-bookmarks", middleware.ClerkAuthMiddleware())
	{
		protected.GET("/:lessonId", ctrl.List)
		protected.POST("", ctrl.Create)
		protected.PATCH("/:id", ctrl.UpdateNote)
		protected.DELETE("/:id", ctrl.Delete)
	}
}
