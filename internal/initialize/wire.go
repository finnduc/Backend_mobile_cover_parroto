package initialize

import (
	"go-familytree/global"
	"go-familytree/internal/controller"
	"go-familytree/internal/repo"
	"go-familytree/internal/service"
)

type AppDependencies struct {
	AuthCtrl     *controller.AuthController
	LessonCtrl   *controller.LessonController
	AttemptCtrl  *controller.AttemptController
	AnswerCtrl   *controller.AnswerController
	ProgressCtrl *controller.ProgressController
	BookmarkCtrl *controller.BookmarkController
	CategoryCtrl *controller.CategoryController
}

func WireDependencies() *AppDependencies {
	// Repos
	authRepo := repo.NewAuthRepo(global.DB)
	lessonRepo := repo.NewLessonRepo(global.DB, global.Rdb)
	attemptRepo := repo.NewAttemptRepo(global.DB)
	answerRepo := repo.NewAnswerRepo(global.DB)
	progressRepo := repo.NewProgressRepo(global.DB)
	bookmarkRepo := repo.NewBookmarkRepo(global.DB)
	categoryRepo := repo.NewCategoryRepo(global.DB)

	// Services
	progressSvc := service.NewProgressService(progressRepo)
	authSvc := service.NewAuthService(authRepo, global.Rdb)
	lessonSvc := service.NewLessonService(lessonRepo)
	attemptSvc := service.NewAttemptService(attemptRepo, lessonRepo)
	answerSvc := service.NewAnswerService(answerRepo, attemptRepo, lessonRepo, progressSvc, global.Config.Scoring.Threshold)
	bookmarkSvc := service.NewBookmarkService(bookmarkRepo)
	categorySvc := service.NewCategoryService(categoryRepo)

	// Controllers
	return &AppDependencies{
		AuthCtrl:     controller.NewAuthController(authSvc),
		LessonCtrl:   controller.NewLessonController(lessonSvc),
		AttemptCtrl:  controller.NewAttemptController(attemptSvc),
		AnswerCtrl:   controller.NewAnswerController(answerSvc),
		ProgressCtrl: controller.NewProgressController(progressSvc),
		BookmarkCtrl: controller.NewBookmarkController(bookmarkSvc),
		CategoryCtrl: controller.NewCategoryController(categorySvc),
	}
}
