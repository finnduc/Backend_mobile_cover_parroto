package bookmark

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.RouterGroup) {
	ctrl := &BookmarkController{}
	r.POST("/bookmarks", ctrl.Add)
	r.DELETE("/bookmarks/:lessonId", ctrl.Remove)
}
