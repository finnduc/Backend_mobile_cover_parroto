package pronunciation

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"go-cover-parroto/internal/middleware"
	"go-cover-parroto/internal/modules/pronunciation/repositories"
	"go-cover-parroto/internal/modules/pronunciation/services"
)

func RegisterRoutes(r *gin.RouterGroup, db *gorm.DB) {
	attemptRepo := repositories.NewPronunciationAttemptRepo(db)
	progressRepo := repositories.NewPronunciationProgressRepo(db)
	svc := services.NewPronunciationService(attemptRepo, progressRepo)
	ctrl := NewPronunciationController(svc)

	protected := r.Group("", middleware.ClerkAuthMiddleware())
	{
		protected.POST("/pronunciation-attempts", ctrl.Assess)
		protected.DELETE("/pronunciation/attempts/:attemptId", ctrl.DeleteAttempt)
		protected.GET("/pronunciation/progress/:lessonId", ctrl.ListProgress)
		protected.GET("/pronunciation/progress/:lessonId/detail", ctrl.ListProgressDetail)
		protected.POST("/pronunciation/progress/update/:transcriptId", ctrl.UpdateProgress)
	}
}
