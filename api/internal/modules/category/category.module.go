package category

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"go-cover-parroto/internal/core/enums"
	"go-cover-parroto/internal/firebase"
	"go-cover-parroto/internal/middleware"
	"go-cover-parroto/internal/modules/category/repositories"
	"go-cover-parroto/internal/modules/category/services"
)

func RegisterRoutes(r *gin.RouterGroup, db *gorm.DB, fbAuth firebase.IFirebaseAuth) {
	repo := repositories.NewCategoryRepo(db)
	svc := services.NewCategoryService(repo)
	ctrl := NewCategoryController(svc)
	adminCtrl := NewCategoryAdminController(svc)

	r.GET("/categories", ctrl.List)

	admin := r.Group("/admin/categories", middleware.FirebaseAuth(fbAuth), middleware.RequireRole(enums.UserRoleAdmin))
	{
		admin.GET("", adminCtrl.List)
		admin.GET("/:id", adminCtrl.GetByID)
		admin.POST("", adminCtrl.Create)
		admin.PUT("/:id", adminCtrl.Update)
		admin.DELETE("/:id", adminCtrl.Delete)
	}
}
