package transcript

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"go-cover-parroto/internal/core/enums"
	"go-cover-parroto/internal/middleware"
	"go-cover-parroto/internal/modules/transcript/repositories"
	"go-cover-parroto/internal/modules/transcript/services"
)

func RegisterRoutes(r *gin.RouterGroup, db *gorm.DB) {
	repo := repositories.NewTranscriptRepo(db)
	svc := services.NewTranscriptService(repo)
	ctrl := NewTranscriptController(svc)
	adminCtrl := NewTranscriptAdminController(svc)

	r.GET("/lessons/:lessonId/transcripts", ctrl.GetByLesson)
	r.GET("/admin/lessons/:lessonId/transcripts", middleware.ClerkAuthMiddleware(), middleware.RequireRole(enums.UserRoleAdmin), adminCtrl.GetByLesson)
	admin := r.Group("/admin/transcripts", middleware.ClerkAuthMiddleware(), middleware.RequireRole(enums.UserRoleAdmin))
	{
		admin.GET("/:id", adminCtrl.GetByID)
		admin.POST("", adminCtrl.Create)
		admin.PUT("/:id", adminCtrl.Update)
		admin.DELETE("/:id", adminCtrl.Delete)
	}
}
