package transcript_progress

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"go-cover-parroto/internal/middleware"
	"go-cover-parroto/internal/modules/transcript_progress/repositories"
	"go-cover-parroto/internal/modules/transcript_progress/services"
)

func RegisterRoutes(r *gin.RouterGroup, db *gorm.DB) {
	repo := repositories.NewTranscriptProgressRepo(db)
	svc := services.NewTranscriptProgressService(repo)
	ctrl := NewTranscriptProgressController(svc)

	protected := r.Group("/transcript-progress", middleware.ClerkAuthMiddleware())
	{
		protected.GET("/:lessonId", ctrl.List)
		protected.POST("/:lessonId", ctrl.Create)
	}
}
