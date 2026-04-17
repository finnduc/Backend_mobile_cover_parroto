package transcript

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"go-cover-parroto/internal/middleware"
	"go-cover-parroto/internal/modules/transcript/repositories"
	"go-cover-parroto/internal/modules/transcript/services"
)

func RegisterRoutes(r *gin.RouterGroup, db *gorm.DB) {
	repo := repositories.NewTranscriptRepo(db)
	svc := services.NewTranscriptService(repo)
	ctrl := NewTranscriptController(svc)

	protected := r.Group("", middleware.FirebaseAuth(db))
	{
		protected.GET("/lessons/:lessonId/transcripts", ctrl.GetByLesson)
	}
}
