package repo

import (
	"context"
	"go-familytree/internal/models"
	"go-familytree/pkg/utils"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type answerRepo struct{ db *gorm.DB }

func NewAnswerRepo(db *gorm.DB) IAnswerRepo {
	return &answerRepo{db: db}
}

func (r *answerRepo) Upsert(ctx context.Context, answer *models.UserAnswer) error {
	return r.db.WithContext(ctx).
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "attempt_id"}, {Name: "transcript_id"}},
			DoUpdates: clause.AssignmentColumns([]string{"answer_text", "score", "is_correct"}),
		}).
		Create(answer).Error
}

func (r *answerRepo) BulkUpsert(ctx context.Context, answers []models.UserAnswer) error {
	if len(answers) == 0 {
		return nil
	}
	return r.db.WithContext(ctx).
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "attempt_id"}, {Name: "transcript_id"}},
			DoUpdates: clause.AssignmentColumns([]string{"answer_text", "score", "is_correct"}),
		}).
		Create(&answers).Error
}

func (r *answerRepo) FindByAttemptAndTranscript(ctx context.Context, attemptID, transcriptID uint) (*models.UserAnswer, error) {
	var ans models.UserAnswer
	err := r.db.WithContext(ctx).
		Where("attempt_id = ? AND transcript_id = ?", attemptID, transcriptID).
		First(&ans).Error
	return &ans, err
}

// ensure utils is used
var _ = utils.PaginationQuery{}
