package transcript

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"go-cover-parroto/internal/core/enums"
	"go-cover-parroto/internal/firebase"
	"go-cover-parroto/internal/middleware"
	"go-cover-parroto/internal/modules/transcript/repositories"
	"go-cover-parroto/internal/modules/transcript/services"
)

func RegisterRoutes(r *gin.RouterGroup, db *gorm.DB, fbAuth firebase.IFirebaseAuth) {
	repo := repositories.NewTranscriptRepo(db)
	svc := services.NewTranscriptService(repo)
	ctrl := NewTranscriptController(svc)
	adminCtrl := NewTranscriptAdminController(svc)

	protected := r.Group("", middleware.FirebaseAuth(db, fbAuth))
	{
		protected.GET("/lessons/:lessonId/transcripts", ctrl.GetByLesson)
	}

	admin := r.Group("/admin", middleware.FirebaseAuth(db, fbAuth), middleware.RequireRole(enums.UserRoleAdmin))
	{
		admin.GET("/transcripts/:id", adminCtrl.GetByID)
		admin.POST("/transcripts", adminCtrl.Create)
		admin.PUT("/transcripts/:id", adminCtrl.Update)
		admin.DELETE("/transcripts/:id", adminCtrl.Delete)
	}
}
