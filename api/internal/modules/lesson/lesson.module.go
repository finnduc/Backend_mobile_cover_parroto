package lesson

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"go-cover-parroto/internal/modules/lesson/repositories"
	"go-cover-parroto/internal/modules/lesson/services"
)

func RegisterRoutes(r *gin.RouterGroup, db *gorm.DB) {
	repo := repositories.NewLessonRepo(db)
	svc := services.NewLessonService(repo)
	ctrl := NewLessonController(svc)

	r.GET("/lessons", ctrl.List)
	r.GET("/lessons/:id", ctrl.Get)
}
