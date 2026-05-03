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

	protected := r.Group("/bookmarks", middleware.FirebaseAuth(fbAuth))
	{
		protected.GET("", ctrl.List)
		protected.POST("/:lessonId", ctrl.Add)
		protected.DELETE("/:lessonId", ctrl.Remove)
	}
}
