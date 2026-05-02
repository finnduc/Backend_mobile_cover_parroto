package lesson

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"go-cover-parroto/internal/firebase"
	"go-cover-parroto/internal/middleware"
	"go-cover-parroto/internal/modules/lesson/repositories"
	"go-cover-parroto/internal/modules/lesson/services"
)

func RegisterRoutes(r *gin.RouterGroup, db *gorm.DB, fbAuth firebase.IFirebaseAuth) {
	repo := repositories.NewLessonRepo(db)
	svc := services.NewLessonService(repo)
	ctrl := NewLessonController(svc)
	adminCtrl := NewLessonAdminController(svc)

	r.GET("/lessons", ctrl.List)
	r.GET("/lessons/:lessonId", ctrl.Get)

	admin := r.Group("/admin", middleware.FirebaseAuth(db, fbAuth))
	admin.POST("/lessons", adminCtrl.Create)
	admin.PUT("/lessons/:id", adminCtrl.Update)
	admin.DELETE("/lessons/:id", adminCtrl.Delete)
}
