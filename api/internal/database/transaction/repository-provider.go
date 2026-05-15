package transaction

import (
	"gorm.io/gorm"

	db_repos "go-cover-parroto/internal/database/repositories"
	authrepos "go-cover-parroto/internal/modules/auth/repositories"
	bookmarkrepos "go-cover-parroto/internal/modules/bookmark/repositories"
	categoryrepos "go-cover-parroto/internal/modules/category/repositories"
	learninghistoryrepos "go-cover-parroto/internal/modules/learning_history/repositories"
	lessonrepos "go-cover-parroto/internal/modules/lesson/repositories"
	transcriptrepos "go-cover-parroto/internal/modules/transcript/repositories"
)

type IProvider interface {
	Auth() db_repos.IAuthRepo
	Bookmark() db_repos.IBookmarkRepo
	Category() db_repos.ICategoryRepo
	LearningHistory() db_repos.ILearningHistoryRepo
	Lesson() db_repos.ILessonRepo
	Transcript() db_repos.ITranscriptRepo
}

type gormProvider struct {
	tx             *gorm.DB
	authRepo       db_repos.IAuthRepo
	bookmarkRepo   db_repos.IBookmarkRepo
	categoryRepo   db_repos.ICategoryRepo
	historyRepo    db_repos.ILearningHistoryRepo
	lessonRepo     db_repos.ILessonRepo
	transcriptRepo db_repos.ITranscriptRepo
}

func NewGormProvider(tx *gorm.DB) *gormProvider {
	return &gormProvider{tx: tx}
}

func (p *gormProvider) Auth() db_repos.IAuthRepo {
	if p.authRepo == nil {
		p.authRepo = authrepos.NewAuthRepo(p.tx)
	}
	return p.authRepo
}

func (p *gormProvider) Bookmark() db_repos.IBookmarkRepo {
	if p.bookmarkRepo == nil {
		p.bookmarkRepo = bookmarkrepos.NewBookmarkRepo(p.tx)
	}
	return p.bookmarkRepo
}

func (p *gormProvider) Category() db_repos.ICategoryRepo {
	if p.categoryRepo == nil {
		p.categoryRepo = categoryrepos.NewCategoryRepo(p.tx)
	}
	return p.categoryRepo
}

func (p *gormProvider) LearningHistory() db_repos.ILearningHistoryRepo {
	if p.historyRepo == nil {
		p.historyRepo = learninghistoryrepos.NewLearningHistoryRepo(p.tx)
	}
	return p.historyRepo
}

func (p *gormProvider) Lesson() db_repos.ILessonRepo {
	if p.lessonRepo == nil {
		p.lessonRepo = lessonrepos.NewLessonRepo(p.tx)
	}
	return p.lessonRepo
}

func (p *gormProvider) Transcript() db_repos.ITranscriptRepo {
	if p.transcriptRepo == nil {
		p.transcriptRepo = transcriptrepos.NewTranscriptRepo(p.tx)
	}
	return p.transcriptRepo
}
