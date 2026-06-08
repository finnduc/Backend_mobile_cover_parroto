package chat

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"go-cover-parroto/internal/middleware"
	"go-cover-parroto/internal/modules/chat/hub"
	"go-cover-parroto/internal/modules/chat/repositories"
	"go-cover-parroto/internal/modules/chat/services"
)

func RegisterRoutes(r *gin.RouterGroup, db *gorm.DB) {
	repo := repositories.NewChatRepo(db)

	sse := hub.NewSSEHub()

	svc := services.NewChatService(repo, sse)
	ctrl := NewChatController(svc, sse)

	protected := r.Group("/chat", middleware.ClerkAuthMiddleware())
	{
		protected.GET("/messages", ctrl.GetHistory)
		protected.POST("/messages", ctrl.SendMessage)
		protected.GET("/events", func(c *gin.Context) {
			sse.ServeHTTP(c.Writer, c.Request)
		})
	}
}
