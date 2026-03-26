package admin

import (
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

func (pr *UserRouter) InitUserRouter(r *gin.RouterGroup) {
	// private
	userPrivate := r.Group("/admin/user")
	// userPrivate.Use(Limit())
	// userPrivate.Use(Auth())
	// userPrivate.Use(Pemission())
	{
		userPrivate.GET("/profile")
		userPrivate.GET("/list")
		userPrivate.GET("/detail/:id")
		userPrivate.POST("/create")
		userPrivate.PUT("/update/:id")
		userPrivate.DELETE("/delete/:id")
	}

}