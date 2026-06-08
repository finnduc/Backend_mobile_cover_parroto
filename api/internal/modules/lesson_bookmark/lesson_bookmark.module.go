package lesson_bookmark

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"go-cover-parroto/internal/middleware"
	"go-cover-parroto/internal/modules/lesson_bookmark/repositories"
	"go-cover-parroto/internal/modules/lesson_bookmark/services"
)

func RegisterRoutes(r *gin.RouterGroup, db *gorm.DB) {
	repo := repositories.NewLessonBookmarkRepo(db)
	svc := services.NewLessonBookmarkService(repo)
	ctrl := NewLessonBookmarkController(svc)

	protected := r.Group("/lesson-bookmarks", middleware.ClerkAuthMiddleware())
	{
		protected.GET("", ctrl.List)
		protected.POST("/:lessonId/toggle", ctrl.Toggle)
		protected.DELETE("/:id", ctrl.Delete)
	}
}
