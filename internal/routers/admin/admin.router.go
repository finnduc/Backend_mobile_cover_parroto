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

	adminPrivate := r.Group("/admin/user")
	// adminPrivate.Use(Limit())
	// adminPrivate.Use(Auth())
	// adminPrivate.Use(Pemission())
	{
		adminPrivate.GET("/list")
		adminPrivate.GET("/detail/:id")
		adminPrivate.POST("/create")
		adminPrivate.PUT("/update/:id")
		adminPrivate.DELETE("/delete/:id")
	}

}
