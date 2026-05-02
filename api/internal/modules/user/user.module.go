package user

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"go-cover-parroto/internal/core/enums"
	"go-cover-parroto/internal/firebase"
	"go-cover-parroto/internal/middleware"
	"go-cover-parroto/internal/modules/user/repositories"
	"go-cover-parroto/internal/modules/user/services"
)

func RegisterRoutes(r *gin.RouterGroup, db *gorm.DB, fbAuth firebase.IFirebaseAuth) {
	repo := repositories.NewUserRepo(db)
	svc := services.NewUserService(repo)
	ctrl := NewUserController(svc)
	adminCtrl := NewUserAdminController(svc)

	protected := r.Group("", middleware.FirebaseAuth(db, fbAuth))
	{
		protected.GET("/user/profile", ctrl.GetProfile)
	}

	admin := r.Group("/admin", middleware.FirebaseAuth(db, fbAuth), middleware.RequireRole(enums.UserRoleAdmin))
	{
		admin.GET("/users", adminCtrl.List)
		admin.GET("/users/:id", adminCtrl.GetByID)
		admin.PUT("/users/:id", adminCtrl.Update)
		admin.DELETE("/users/:id", adminCtrl.Delete)
	}
}
