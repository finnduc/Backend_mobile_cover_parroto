package learning_history

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"go-cover-parroto/internal/middleware"
	"go-cover-parroto/internal/modules/learning_history/repositories"
	"go-cover-parroto/internal/modules/learning_history/services"
)

func RegisterRoutes(r *gin.RouterGroup, db *gorm.DB) {
	repo := repositories.NewLearningHistoryRepo(db)
	svc := services.NewLearningHistoryService(repo)
	ctrl := NewLearningHistoryController(svc)

	protected := r.Group("/learning-history", middleware.ClerkAuthMiddleware())
	{
		protected.GET("", ctrl.List)
		protected.POST("", ctrl.Create)
		protected.GET("/finished", ctrl.ListFinished)
		protected.GET("/unfinished", ctrl.ListUnfinished)
		protected.GET("/summary", ctrl.Summary)
		protected.GET("/lessons/:lessonId/summary", ctrl.LessonSummary)
		protected.GET("/:lessonId", ctrl.GetByLesson)
	}
}
