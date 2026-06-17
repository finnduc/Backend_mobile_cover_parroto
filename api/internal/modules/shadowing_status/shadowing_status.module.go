package shadowing_status

import (
	"go-cover-parroto/internal/middleware"
	"go-cover-parroto/internal/modules/shadowing_status/repositories"
	"go-cover-parroto/internal/modules/shadowing_status/services"

	deepgram "github.com/deepgram/deepgram-go-sdk/v3/pkg/api/listen/v1/rest"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.RouterGroup, db *gorm.DB, dgApi *deepgram.Client) {
	repo := repositories.NewShadowingStatusRepo(db)
	svc := services.NewShadowingStatusService(repo)
	transcriptionSvc := services.NewTranscriptionService(dgApi)
	ctrl := NewShadowingStatusController(svc, transcriptionSvc)

	protected := r.Group("/shadowing-status", middleware.ClerkAuthMiddleware())
	{
		protected.POST("", ctrl.Create)
		protected.GET("", ctrl.List)
		protected.POST("/transcribe", ctrl.TranscribeShadowing)
	}
}
