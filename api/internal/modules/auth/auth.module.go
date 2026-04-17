package auth

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"go-cover-parroto/internal/modules/auth/repositories"
	"go-cover-parroto/internal/modules/auth/services"
)

func RegisterRoutes(r *gin.RouterGroup, db *gorm.DB) {
	repo := repositories.NewAuthRepo(db)
	svc := services.NewAuthService(repo)
	ctrl := NewAuthController(svc)

	r.POST("/auth/sync", ctrl.Sync)
}
