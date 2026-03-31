package admin

import (
	"github.com/gin-gonic/gin"
)

type AdminRouter struct {
}

func (pr *AdminRouter) InitAdminRouter(r *gin.RouterGroup) {

	adminPublic := r.Group("/admin")
	{
		adminPublic.GET("/login")
	}
}
