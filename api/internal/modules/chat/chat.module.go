package chat

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"go-cover-parroto/internal/middleware"
	"go-cover-parroto/internal/modules/chat/hub"
	"go-cover-parroto/internal/modules/chat/repositories"
	"go-cover-parroto/internal/modules/chat/services"
	userrepo "go-cover-parroto/internal/modules/user/repositories"
)

func RegisterRoutes(r *gin.RouterGroup, db *gorm.DB) {
	repo := repositories.NewChatRepo(db)
	userRepo := userrepo.NewUserRepo(db)

	h := hub.NewHub()
	go h.Run()

	svc := services.NewChatService(repo, userRepo, h)
	ctrl := NewChatController(svc, h)

	protected := r.Group("/chat", middleware.ClerkAuthMiddleware())
	{
		protected.GET("/messages", ctrl.GetHistory)
		protected.GET("/ws", ctrl.Connect)
	}
}
