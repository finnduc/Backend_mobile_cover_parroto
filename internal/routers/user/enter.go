package user

import (
	"go-familytree/internal/controller"

	"github.com/gin-gonic/gin"
)


type UserRouterGroup struct {
	UserRouter     UserRouter
	LessonRouter   LessonRouter
	CategoryRouter CategoryRouter
	ProgressRouter ProgressRouter
	BookmarkRouter BookmarkRouter
}

func (pr *UserRouterGroup) InitUserRouterGroup(r *gin.RouterGroup, 
	authCtrl *controller.AuthController,
	lessonCtrl *controller.LessonController,
	attemptCtrl *controller.AttemptController,
	answerCtrl *controller.AnswerController,
	progressCtrl *controller.ProgressController,
	bookmarkCtrl *controller.BookmarkController,
	categoryCtrl *controller.CategoryController,
) {
	pr.UserRouter.InitUserRouter(r, authCtrl)
	pr.LessonRouter.InitLessonRouter(r, lessonCtrl, attemptCtrl, answerCtrl)
	pr.CategoryRouter.InitCategoryRouter(r, categoryCtrl)
	pr.ProgressRouter.InitProgressRouter(r, progressCtrl)
	pr.BookmarkRouter.InitBookmarkRouter(r, bookmarkCtrl)
}
