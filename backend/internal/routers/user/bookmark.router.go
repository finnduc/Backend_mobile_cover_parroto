package user

import (
	"go-cover-parroto/internal/controller"
	"go-cover-parroto/internal/middlewares"

	"github.com/gin-gonic/gin"
)

type BookmarkRouter struct {
}

func (pr *BookmarkRouter) InitBookmarkRouter(r *gin.RouterGroup, bookmarkCtrl *controller.BookmarkController) {
	bookmark := r.Group("/bookmarks")
	bookmark.Use(middlewares.AuthMiddleware())
	{
		bookmark.POST("/", bookmarkCtrl.Toggle)
		bookmark.GET("/", bookmarkCtrl.List)
	}
}
