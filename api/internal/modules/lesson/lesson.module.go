package lesson

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"go-cover-parroto/internal/core/enums"
	"go-cover-parroto/internal/middleware"
	"go-cover-parroto/internal/modules/lesson/repositories"
	"go-cover-parroto/internal/modules/lesson/services"
)

func RegisterRoutes(r *gin.RouterGroup, db *gorm.DB) {
	repo := repositories.NewLessonRepo(db)
	svc := services.NewLessonService(repo)
	ctrl := NewLessonController(svc)
	adminCtrl := NewLessonAdminController(svc)

	r.GET("/lessons", ctrl.List)
	r.GET("/lessons/:lessonId", ctrl.Get)

	admin := r.Group("/admin/lessons", middleware.ClerkAuthMiddleware(), middleware.RequireRole(enums.UserRoleAdmin))
	{
		admin.GET("", adminCtrl.List)
		admin.GET("/:id", adminCtrl.Get)
		admin.POST("", adminCtrl.Create)
		admin.PUT("/:id", adminCtrl.Update)
		admin.DELETE("/:id", adminCtrl.Delete)
	}
}
