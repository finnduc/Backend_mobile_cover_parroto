package services

import (
	"context"
	"time"

	"go-cover-parroto/internal/database/models"
	"go-cover-parroto/internal/modules/lesson/dtos/res"
	"go-cover-parroto/internal/modules/lesson/repositories"

	"gorm.io/gorm"
)

var lessonRepo repositories.ILessonRepo

func ListLessons() ([]res.LessonRes, error) {
	lessons, err := lessonRepo.FindAll(context.Background())
	if err != nil {
		return nil, err
	}
	var result []res.LessonRes
	for _, lesson := range lessons {
		result = append(result, toLessonRes(lesson))
	}
	return result, nil
}

func GetLesson(id uint) (*res.LessonRes, error) {
	lesson, err := lessonRepo.FindByID(context.Background(), id)
	if err != nil {
		return nil, err
	}
	res := toLessonRes(*lesson)
	return &res, nil
}

func toLessonRes(lesson models.Lesson) res.LessonRes {
	return res.LessonRes{
		ID:           lesson.ID,
		CategoryID:   lesson.CategoryID,
		Title:        lesson.Title,
		Description:  lesson.Description,
		VideoURL:     lesson.VideoURL,
		ThumbnailURL: lesson.ThumbnailURL,
		Level:        lesson.Level,
		Duration:     lesson.Duration,
		CreatedAt:    lesson.CreatedAt.Format(time.RFC3339),
	}
}

func SetLessonRepo(db *gorm.DB) {
	lessonRepo = repositories.NewLessonRepo(db)
}
