package auth

import (
	"github.com/gin-gonic/gin"

	"go-cover-parroto/internal/middleware"
	"go-cover-parroto/internal/modules/auth/services"
)

func RegisterRoutes(r *gin.RouterGroup) {
	svc := services.NewAuthService()
	ctrl := NewAuthController(svc)

	protected := r.Group("/auth", middleware.ClerkAuthMiddleware())
	{
		protected.POST("/complete-signup", ctrl.Complete)
	}
}
