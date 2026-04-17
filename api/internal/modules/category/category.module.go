package category

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.RouterGroup) {
	ctrl := &CategoryController{}
	r.GET("/categories", ctrl.List)
}
