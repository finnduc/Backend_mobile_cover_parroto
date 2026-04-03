package admin

import "github.com/gin-gonic/gin"

type AdminRouterGroup struct {
	AdminRouter
	UserRouter
}

func (pr *AdminRouterGroup) InitAdminRouterGroup(r *gin.RouterGroup) {
	pr.AdminRouter.InitAdminRouter(r)
	pr.UserRouter.InitUserRouter(r)
}