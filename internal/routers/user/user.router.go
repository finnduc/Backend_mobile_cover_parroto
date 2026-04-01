package user

import (
	"go-cover-parroto/internal/controller"
	"go-cover-parroto/internal/middlewares"

	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

func (pr *UserRouter) InitUserRouter(r *gin.RouterGroup, authCtrl *controller.AuthController) {
	auth := r.Group("/auth")
	{
		auth.POST("/register", authCtrl.Register)
		
		auth.POST("/login", authCtrl.Login)
		auth.POST("/refresh", authCtrl.Refresh)
		auth.POST("/logout", middlewares.AuthMiddleware(), authCtrl.Logout)
	}

	// User group
	userGroup := r.Group("/user")
	userGroup.Use(middlewares.AuthMiddleware())
	{
		userGroup.GET("/profile", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "profile"})
		})
	}
}