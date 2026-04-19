package bookmark

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"go-cover-parroto/internal/firebase"
	"go-cover-parroto/internal/middleware"
	"go-cover-parroto/internal/modules/bookmark/repositories"
	"go-cover-parroto/internal/modules/bookmark/services"
)

func RegisterRoutes(r *gin.RouterGroup, db *gorm.DB, fbAuth firebase.IFirebaseAuth) {
	repo := repositories.NewBookmarkRepo(db)
	svc := services.NewBookmarkService(repo)
	ctrl := NewBookmarkController(svc)

	protected := r.Group("", middleware.FirebaseAuth(db, fbAuth))
	{
		protected.GET("/bookmarks", ctrl.List)
		protected.POST("/bookmarks/:lessonId", ctrl.Add)
		protected.DELETE("/bookmarks/:lessonId", ctrl.Remove)
	}
}
