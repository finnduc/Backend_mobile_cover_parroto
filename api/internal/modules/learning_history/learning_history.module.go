package learning_history

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"go-cover-parroto/internal/firebase"
	"go-cover-parroto/internal/middleware"
	"go-cover-parroto/internal/modules/learning_history/repositories"
	"go-cover-parroto/internal/modules/learning_history/services"
)

func RegisterRoutes(r *gin.RouterGroup, db *gorm.DB, fbAuth firebase.IFirebaseAuth) {
	repo := repositories.NewLearningHistoryRepo(db)
	svc := services.NewLearningHistoryService(repo)
	ctrl := NewLearningHistoryController(svc)

	protected := r.Group("", middleware.FirebaseAuth(db, fbAuth))
	{
		protected.POST("/learning-history", ctrl.Record)
		protected.GET("/learning-history", ctrl.List)
		protected.GET("/learning-history/:lessonId", ctrl.GetByLesson)
	}
}
