package dictation_status

import (
	"go-cover-parroto/internal/middleware"
	"go-cover-parroto/internal/modules/dictation_status/repositories"
	"go-cover-parroto/internal/modules/dictation_status/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.RouterGroup, db *gorm.DB) {
	repo := repositories.NewDictationStatusRepo(db)
	svc := services.NewDictationStatusService(repo)
	ctrl := NewDictationStatusController(svc)

	protected := r.Group("/dictation-status", middleware.ClerkAuthMiddleware())
	{
		protected.POST("", ctrl.Create)
		protected.GET("", ctrl.List)
	}
}
