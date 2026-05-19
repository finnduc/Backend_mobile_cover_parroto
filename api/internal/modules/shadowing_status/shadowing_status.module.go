package shadowing_status

import (
	"go-cover-parroto/internal/middleware"
	"go-cover-parroto/internal/modules/shadowing_status/repositories"
	"go-cover-parroto/internal/modules/shadowing_status/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.RouterGroup, db *gorm.DB) {
	repo := repositories.NewShadowingStatusRepo(db)
	svc := services.NewShadowingStatusService(repo)
	ctrl := NewShadowingStatusController(svc)

	protected := r.Group("/shadowing-status", middleware.ClerkAuthMiddleware())
	{
		protected.POST("/:transcriptId", ctrl.Create)
		protected.GET("", ctrl.List)
	}
}
