package routers

import (
	"go-familytree/internal/controller"
	"go-familytree/internal/routers/admin"
	"go-familytree/internal/routers/user"

	"github.com/gin-gonic/gin"
)

type RouterGroup struct {
	UserRouterGroup  user.UserRouterGroup
	AdminRouterGroup admin.AdminRouterGroup
}

var RouterGroupApp = new(RouterGroup)

func (pr *RouterGroup) InitRouterGroup(r *gin.RouterGroup,
	authCtrl *controller.AuthController,
	lessonCtrl *controller.LessonController,
	attemptCtrl *controller.AttemptController,
	answerCtrl *controller.AnswerController,
	progressCtrl *controller.ProgressController,
	bookmarkCtrl *controller.BookmarkController,
	categoryCtrl *controller.CategoryController,
) {
	pr.UserRouterGroup.InitUserRouterGroup(r, authCtrl, lessonCtrl, attemptCtrl, answerCtrl, progressCtrl, bookmarkCtrl, categoryCtrl)
	pr.AdminRouterGroup.InitAdminRouterGroup(r)
}