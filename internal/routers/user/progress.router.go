package user

import (
	"go-familytree/internal/controller"
	"go-familytree/internal/middlewares"

	"github.com/gin-gonic/gin"
)

type ProgressRouter struct {
}

func (pr *ProgressRouter) InitProgressRouter(r *gin.RouterGroup, progressCtrl *controller.ProgressController) {
	progress := r.Group("/progress")
	progress.Use(middlewares.AuthMiddleware())
	{
		progress.GET("/:lesson_id", progressCtrl.Get)
	}
}
