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
	repo := repositories.NewUserRepo(fbAuth)
	svc := services.NewUserService(repo)
	ctrl := NewUserController(svc)
	adminCtrl := NewUserAdminController(svc)

	protected := r.Group("", middleware.FirebaseAuth(fbAuth))
	{
		protected.GET("/user/profile", ctrl.GetProfile)
	}

	admin := r.Group("/admin/users", middleware.FirebaseAuth(fbAuth), middleware.RequireRole(enums.UserRoleAdmin))
	{
		admin.GET("", adminCtrl.List)
		admin.GET("/:id", adminCtrl.GetByID)
		// admin.PUT("/:id", adminCtrl.Update)
		admin.DELETE("/:id", adminCtrl.Delete)
	}
}
