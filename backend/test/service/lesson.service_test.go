package service

import (
	"context"
	"go-cover-parroto/internal/mocks"
	"go-cover-parroto/internal/service"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateLesson(t *testing.T) {

	mockrepo := new(mocks.MockLessonRepo)

	svc := service.NewLessonService(mockrepo)
		

	input := service.CreateLessonInput{
		Title: "Test Lesson",
		Description: "Test Description",
		VideoURL: "https://www.youtube.com/watch?v=dQw4w9WgXcQ",
		ThumbnailURL: "https://www.youtube.com/watch?v=dQw4w9WgXcQ",
	}

	mockrepo.On("Create",
		mock.Anything,
		mock.AnythingOfType("*models.Lesson"),
		mock.AnythingOfType("[]uint"),
		mock.AnythingOfType("[]models.Transcript"),
	).Return(nil)

	lesson, err := svc.CreateLesson(context.Background(), input)
	
	assert.NoError(t, err)
	assert.NotNil(t, lesson)

	mockrepo.AssertExpectations(t)
}