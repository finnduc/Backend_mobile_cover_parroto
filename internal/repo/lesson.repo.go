package repo

import (
	"context"
	"encoding/json"
	"fmt"
	"go-familytree/internal/models"
	"go-familytree/pkg/utils"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type lessonRepo struct {
	db  *gorm.DB
	rdb *redis.Client
}

func NewLessonRepo(db *gorm.DB, rdb *redis.Client) ILessonRepo {
	return &lessonRepo{db: db, rdb: rdb}
}

func (r *lessonRepo) FindAll(ctx context.Context, filter LessonFilter) ([]models.Lesson, int64, error) {
	filter.Normalize()
	tx := r.db.WithContext(ctx).Model(&models.Lesson{})
	if filter.Level != "" {
		tx = tx.Where("level = ?", filter.Level)
	}
	if filter.CategoryID != nil {
		tx = tx.Joins("JOIN lesson_categories lc ON lc.lesson_id = lessons.id").
			Where("lc.category_id = ?", *filter.CategoryID)
	}

	var total int64
	if err := tx.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var lessons []models.Lesson
	err := tx.Offset(filter.Offset()).Limit(filter.Limit).Find(&lessons).Error
	return lessons, total, err
}

func (r *lessonRepo) Create(ctx context.Context, lesson *models.Lesson, categoryIDs []uint, transcripts []models.Transcript) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(lesson).Error; err != nil {
			return err
		}

		if len(categoryIDs) > 0 {
			lessonCategories := make([]models.LessonCategory, 0, len(categoryIDs))
			for _, categoryID := range categoryIDs {
				lessonCategories = append(lessonCategories, models.LessonCategory{
					LessonID:   lesson.ID,
					CategoryID: categoryID,
				})
			}
			if err := tx.Create(&lessonCategories).Error; err != nil {
				return err
			}
		}

		if len(transcripts) > 0 {
			for i := range transcripts {
				transcripts[i].LessonID = lesson.ID
			}
			if err := tx.Create(&transcripts).Error; err != nil {
				return err
			}
		}

		return nil
	})
}

func (r *lessonRepo) FindByID(ctx context.Context, id uint) (*models.Lesson, error) {
	var lesson models.Lesson
	err := r.db.WithContext(ctx).First(&lesson, id).Error
	return &lesson, err
}

// FindTranscriptsByLessonID returns transcripts with Redis caching (TTL 1h)
func (r *lessonRepo) FindTranscriptsByLessonID(ctx context.Context, lessonID uint) ([]models.Transcript, error) {
	cacheKey := fmt.Sprintf("transcripts:%d", lessonID)

	// Try cache first
	if r.rdb != nil {
		cached, err := r.rdb.Get(ctx, cacheKey).Bytes()
		if err == nil {
			var transcripts []models.Transcript
			if err := json.Unmarshal(cached, &transcripts); err == nil {
				return transcripts, nil
			}
		}
	}

	var transcripts []models.Transcript
	err := r.db.WithContext(ctx).
		Where("lesson_id = ?", lessonID).
		Order("sequence ASC").
		Find(&transcripts).Error
	if err != nil {
		return nil, err
	}

	// Store in cache
	if r.rdb != nil {
		if data, err := json.Marshal(transcripts); err == nil {
			r.rdb.Set(ctx, cacheKey, data, time.Hour)
		}
	}
	_ = utils.PaginationQuery{} // ensure no unused import
	return transcripts, nil
}
