package transcript_bookmark

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"go-cover-parroto/internal/middleware"
	transcriptrepos "go-cover-parroto/internal/modules/transcript/repositories"
	"go-cover-parroto/internal/modules/transcript_bookmark/repositories"
	"go-cover-parroto/internal/modules/transcript_bookmark/services"
)

func RegisterRoutes(r *gin.RouterGroup, db *gorm.DB) {
	repo := repositories.NewTranscriptBookmarkRepo(db)
	transcriptRepo := transcriptrepos.NewTranscriptRepo(db)
	svc := services.NewTranscriptBookmarkService(repo, transcriptRepo)
	ctrl := NewTranscriptBookmarkController(svc)

	protected := r.Group("/transcript-bookmarks", middleware.ClerkAuthMiddleware())
	{
		protected.GET("", ctrl.List)
		protected.GET("/:lessonId", ctrl.List)
		protected.POST("", ctrl.Create)
		protected.PUT("/:transcriptId", ctrl.Update)
		protected.DELETE("/:transcriptId", ctrl.Delete)
	}
}
