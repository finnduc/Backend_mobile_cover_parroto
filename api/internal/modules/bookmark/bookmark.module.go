package bookmark

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"go-cover-parroto/internal/middleware"
	"go-cover-parroto/internal/modules/bookmark/repositories"
	"go-cover-parroto/internal/modules/bookmark/services"
)

func RegisterRoutes(r *gin.RouterGroup, db *gorm.DB) {
	repo := repositories.NewBookmarkRepo(db)
	svc := services.NewBookmarkService(repo)
	ctrl := NewBookmarkController(svc)

	protected := r.Group("", middleware.FirebaseAuth(db))
	{
		protected.POST("/bookmarks/:lessonId", ctrl.Add)
		protected.DELETE("/bookmarks/:lessonId", ctrl.Remove)
	}
}
