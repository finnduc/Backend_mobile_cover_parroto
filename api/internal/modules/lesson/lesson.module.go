package lesson

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.RouterGroup) {
	ctrl := &LessonController{}
	r.GET("/lessons", ctrl.List)
	r.GET("/lessons/:id", ctrl.Get)
}
