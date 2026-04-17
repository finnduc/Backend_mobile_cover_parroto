package progress

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.RouterGroup) {
	ctrl := &ProgressController{}
	r.POST("/progress", ctrl.Update)
	r.GET("/progress/:lessonId", ctrl.Get)
}
