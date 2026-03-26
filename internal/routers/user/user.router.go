package user

import (
	"go-familytree/internal/controller"
	"go-familytree/internal/middlewares"
	"go-familytree/internal/wire"

	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

func (pr *UserRouter) InitUserRouter(r *gin.RouterGroup, authCtrl *controller.AuthController) {
	// Use wire to get the new UserController
	userController, _ := wire.InitUserRouterHandle()

	auth := r.Group("/auth")
	{
		// Use userController for register as requested/implied
		auth.POST("/register", userController.Register)
		
		// Use existing authCtrl for other auth actions
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