package user

import (
	"go-familytree/internal/controller"
	"go-familytree/internal/middlewares"

	"github.com/gin-gonic/gin"
)

type LessonRouter struct {
}

func (pr *LessonRouter) InitLessonRouter(r *gin.RouterGroup, 
	lessonCtrl *controller.LessonController,
	attemptCtrl *controller.AttemptController,
	answerCtrl *controller.AnswerController,
) {
	lesson := r.Group("/lessons")
	lesson.Use(middlewares.AuthMiddleware())
	{
		lesson.GET("/", lessonCtrl.ListLessons)
		lesson.GET("/:id", lessonCtrl.GetLesson)
		lesson.GET("/:id/transcripts", lessonCtrl.GetTranscripts)
	}

	// Attempts
	attempts := r.Group("/attempts")
	attempts.Use(middlewares.AuthMiddleware())
	{
		attempts.POST("/", attemptCtrl.Create)
		attempts.GET("/:id", attemptCtrl.Get)
	}

	// Answers
	answers := r.Group("/answers")
	answers.Use(middlewares.AuthMiddleware())
	{
		answers.POST("/", answerCtrl.Submit)
		answers.POST("/bulk", answerCtrl.BulkSubmit)
	}
}