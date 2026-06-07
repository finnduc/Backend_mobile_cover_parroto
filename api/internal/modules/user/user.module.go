package user

import (
	"github.com/gin-gonic/gin"

	"go-cover-parroto/internal/middleware"
	"go-cover-parroto/internal/modules/auth/services"
)

func RegisterRoutes(r *gin.RouterGroup) {
	svc := services.NewAuthService()
	ctrl := NewUserController(svc)

	protected := r.Group("/user", middleware.ClerkAuthMiddleware())
	{
		protected.GET("/profile", ctrl.GetProfile)
	}
}
