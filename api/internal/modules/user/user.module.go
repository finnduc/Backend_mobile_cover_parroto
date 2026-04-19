package user

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"go-cover-parroto/internal/firebase"
	"go-cover-parroto/internal/middleware"
	"go-cover-parroto/internal/modules/user/repositories"
	"go-cover-parroto/internal/modules/user/services"
)

func RegisterRoutes(r *gin.RouterGroup, db *gorm.DB, fbAuth firebase.IFirebaseAuth) {
	repo := repositories.NewUserRepo(db)
	svc := services.NewUserService(repo)
	ctrl := NewUserController(svc)

	protected := r.Group("", middleware.FirebaseAuth(db, fbAuth))
	{
		protected.GET("/user/profile", ctrl.GetProfile)
	}
}
