package services

import (
	"context"
	"errors"
	"net/http"
	"testing"

	coreError "go-cover-parroto/internal/core/errors"
	"go-cover-parroto/internal/core/enums"
	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/database/models"
	"go-cover-parroto/internal/modules/vocabulary_deck/dtos/req"
	"go-cover-parroto/internal/modules/vocabulary_deck/repositories"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func userCtx(userID string) context.Context {
	ctx := context.WithValue(context.Background(), enums.ContextKeyUserID, userID)
	return context.WithValue(ctx, enums.ContextKeyUserRole, enums.UserRoleUser)
}

func adminCtx() context.Context {
	ctx := context.WithValue(context.Background(), enums.ContextKeyUserID, "admin1")
	return context.WithValue(ctx, enums.ContextKeyUserRole, enums.UserRoleAdmin)
}

func strPtr(s string) *string { return &s }

func TestVocabularyDeckService_ListDefault(t *testing.T) {
	tests := []struct {
		name     string
		setup    func(*repositories.MockVocabularyDeckRepo)
		wantLen  int
		wantErr  bool
		wantCode int
	}{
		{
			name: "success — returns default decks",
			setup: func(r *repositories.MockVocabularyDeckRepo) {
				isDefault := true
				r.On("FindAll", mock.Anything, mock.MatchedBy(func(q interface{}) bool { return true })).Return(
					&response.PaginatedResult[*models.VocabularyDeck]{
						Data: []*models.VocabularyDeck{
							{ID: 1, Name: "Basic Greetings", IsDefault: true},
							{ID: 2, Name: "Phrasal Verbs", IsDefault: true},
						},
						Meta: response.NewMeta(1, 10, 2),
					}, nil)
				_ = isDefault
			},
			wantLen: 2,
		},
		{
			name: "db error returns 500",
			setup: func(r *repositories.MockVocabularyDeckRepo) {
				r.On("FindAll", mock.Anything, mock.Anything).Return(nil, errors.New("db error"))
			},
			wantErr:  true,
			wantCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(repositories.MockVocabularyDeckRepo)
			tt.setup(mockRepo)
			svc := NewVocabularyDeckService(mockRepo)

			result, err := svc.ListDefault(context.Background(), req.ListVocabularyDeckQuery{Page: 1, Limit: 10})

			if tt.wantErr {
				assert.NotNil(t, err)
				assert.Equal(t, tt.wantCode, err.Code)
			} else {
				assert.Nil(t, err)
				assert.Len(t, result.Data, tt.wantLen)
			}
			mockRepo.AssertExpectations(t)
		})
	}
}

func TestVocabularyDeckService_ListByUser(t *testing.T) {
	tests := []struct {
		name     string
		userID   string
		setup    func(*repositories.MockVocabularyDeckRepo)
		wantLen  int
		wantErr  bool
		wantCode int
	}{
		{
			name:   "success — returns user decks",
			userID: "user1",
			setup: func(r *repositories.MockVocabularyDeckRepo) {
				r.On("FindAll", mock.Anything, mock.Anything).Return(
					&response.PaginatedResult[*models.VocabularyDeck]{
						Data: []*models.VocabularyDeck{
							{ID: 4, Name: "My Deck", UserID: strPtr("user1")},
						},
						Meta: response.NewMeta(1, 10, 1),
					}, nil)
			},
			wantLen: 1,
		},
		{
			name:   "db error returns 500",
			userID: "user1",
			setup: func(r *repositories.MockVocabularyDeckRepo) {
				r.On("FindAll", mock.Anything, mock.Anything).Return(nil, errors.New("db error"))
			},
			wantErr:  true,
			wantCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(repositories.MockVocabularyDeckRepo)
			tt.setup(mockRepo)
			svc := NewVocabularyDeckService(mockRepo)

			result, err := svc.ListByUser(userCtx("user1"), tt.userID, req.ListVocabularyDeckQuery{Page: 1, Limit: 10})

			if tt.wantErr {
				assert.NotNil(t, err)
				assert.Equal(t, tt.wantCode, err.Code)
			} else {
				assert.Nil(t, err)
				assert.Len(t, result.Data, tt.wantLen)
			}
			mockRepo.AssertExpectations(t)
		})
	}
}

func TestVocabularyDeckService_GetByID(t *testing.T) {
	tests := []struct {
		name     string
		id       uint
		setup    func(*repositories.MockVocabularyDeckRepo)
		wantID   uint
		wantErr  bool
		wantCode int
	}{
		{
			name: "success",
			id:   1,
			setup: func(r *repositories.MockVocabularyDeckRepo) {
				r.On("FindByID", mock.Anything, uint(1)).Return(
					&models.VocabularyDeck{ID: 1, Name: "Basic Greetings"}, nil)
			},
			wantID: 1,
		},
		{
			name: "not found returns 404",
			id:   999,
			setup: func(r *repositories.MockVocabularyDeckRepo) {
				r.On("FindByID", mock.Anything, uint(999)).Return(nil, coreError.ErrNotFound)
			},
			wantErr:  true,
			wantCode: http.StatusNotFound,
		},
		{
			name: "db error returns 500",
			id:   1,
			setup: func(r *repositories.MockVocabularyDeckRepo) {
				r.On("FindByID", mock.Anything, uint(1)).Return(nil, errors.New("db error"))
			},
			wantErr:  true,
			wantCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(repositories.MockVocabularyDeckRepo)
			tt.setup(mockRepo)
			svc := NewVocabularyDeckService(mockRepo)

			result, err := svc.GetByID(context.Background(), tt.id)

			if tt.wantErr {
				assert.NotNil(t, err)
				assert.Equal(t, tt.wantCode, err.Code)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.wantID, result.ID)
			}
			mockRepo.AssertExpectations(t)
		})
	}
}

func TestVocabularyDeckService_Create(t *testing.T) {
	tests := []struct {
		name     string
		userID   string
		body     req.CreateVocabularyDeckReq
		setup    func(*repositories.MockVocabularyDeckRepo)
		wantName string
		wantErr  bool
		wantCode int
	}{
		{
			name:   "success — creates user deck",
			userID: "user1",
			body:   req.CreateVocabularyDeckReq{Name: "My Deck", Level: "beginner"},
			setup: func(r *repositories.MockVocabularyDeckRepo) {
				r.On("Create", mock.Anything, mock.MatchedBy(func(d *models.VocabularyDeck) bool {
					return d.Name == "My Deck" && d.UserID != nil && *d.UserID == "user1"
				})).Run(func(args mock.Arguments) {
					args.Get(1).(*models.VocabularyDeck).ID = 4
				}).Return(nil)
			},
			wantName: "My Deck",
		},
		{
			name:   "db error returns 500",
			userID: "user1",
			body:   req.CreateVocabularyDeckReq{Name: "Fail"},
			setup: func(r *repositories.MockVocabularyDeckRepo) {
				r.On("Create", mock.Anything, mock.Anything).Return(errors.New("db error"))
			},
			wantErr:  true,
			wantCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(repositories.MockVocabularyDeckRepo)
			tt.setup(mockRepo)
			svc := NewVocabularyDeckService(mockRepo)

			result, err := svc.Create(userCtx("user1"), tt.userID, tt.body)

			if tt.wantErr {
				assert.NotNil(t, err)
				assert.Equal(t, tt.wantCode, err.Code)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.wantName, result.Name)
				assert.NotNil(t, result.UserID)
			}
			mockRepo.AssertExpectations(t)
		})
	}
}

func TestVocabularyDeckService_Update(t *testing.T) {
	ownerID := "user1"
	userDeck := &models.VocabularyDeck{ID: 4, Name: "Old Name", UserID: &ownerID}

	tests := []struct {
		name     string
		ctx      context.Context
		id       uint
		body     req.UpdateVocabularyDeckReq
		setup    func(*repositories.MockVocabularyDeckRepo)
		wantName string
		wantErr  bool
		wantCode int
	}{
		{
			name: "success — owner updates deck",
			ctx:  userCtx("user1"),
			id:   4,
			body: req.UpdateVocabularyDeckReq{Name: strPtr("New Name"), Level: strPtr("intermediate")},
			setup: func(r *repositories.MockVocabularyDeckRepo) {
				r.On("FindByID", mock.Anything, uint(4)).Return(userDeck, nil)
				r.On("Update", mock.Anything, mock.MatchedBy(func(d *models.VocabularyDeck) bool {
					return d.Name == "New Name"
				})).Return(nil)
			},
			wantName: "New Name",
		},
		{
			name: "forbidden — non-owner returns 403",
			ctx:  userCtx("other_user"),
			id:   4,
			body: req.UpdateVocabularyDeckReq{Name: strPtr("Hack")},
			setup: func(r *repositories.MockVocabularyDeckRepo) {
				r.On("FindByID", mock.Anything, uint(4)).Return(userDeck, nil)
			},
			wantErr:  true,
			wantCode: http.StatusForbidden,
		},
		{
			name: "not found returns 404",
			ctx:  userCtx("user1"),
			id:   999,
			body: req.UpdateVocabularyDeckReq{Name: strPtr("X")},
			setup: func(r *repositories.MockVocabularyDeckRepo) {
				r.On("FindByID", mock.Anything, uint(999)).Return(nil, coreError.ErrNotFound)
			},
			wantErr:  true,
			wantCode: http.StatusNotFound,
		},
		{
			name: "update error returns 500",
			ctx:  userCtx("user1"),
			id:   4,
			body: req.UpdateVocabularyDeckReq{Name: strPtr("Name")},
			setup: func(r *repositories.MockVocabularyDeckRepo) {
				r.On("FindByID", mock.Anything, uint(4)).Return(userDeck, nil)
				r.On("Update", mock.Anything, mock.Anything).Return(errors.New("db error"))
			},
			wantErr:  true,
			wantCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(repositories.MockVocabularyDeckRepo)
			tt.setup(mockRepo)
			svc := NewVocabularyDeckService(mockRepo)

			result, err := svc.Update(tt.ctx, tt.id, tt.body)

			if tt.wantErr {
				assert.NotNil(t, err)
				assert.Equal(t, tt.wantCode, err.Code)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.wantName, result.Name)
			}
			mockRepo.AssertExpectations(t)
		})
	}
}

func TestVocabularyDeckService_Delete(t *testing.T) {
	ownerID := "user1"
	userDeck := &models.VocabularyDeck{ID: 4, Name: "My Deck", UserID: &ownerID}

	tests := []struct {
		name     string
		ctx      context.Context
		id       uint
		setup    func(*repositories.MockVocabularyDeckRepo)
		wantErr  bool
		wantCode int
	}{
		{
			name: "success — owner deletes deck",
			ctx:  userCtx("user1"),
			id:   4,
			setup: func(r *repositories.MockVocabularyDeckRepo) {
				r.On("FindByID", mock.Anything, uint(4)).Return(userDeck, nil)
				r.On("Delete", mock.Anything, uint(4)).Return(nil)
			},
		},
		{
			name: "forbidden — non-owner returns 403",
			ctx:  userCtx("other_user"),
			id:   4,
			setup: func(r *repositories.MockVocabularyDeckRepo) {
				r.On("FindByID", mock.Anything, uint(4)).Return(userDeck, nil)
			},
			wantErr:  true,
			wantCode: http.StatusForbidden,
		},
		{
			name: "not found returns 404",
			ctx:  userCtx("user1"),
			id:   999,
			setup: func(r *repositories.MockVocabularyDeckRepo) {
				r.On("FindByID", mock.Anything, uint(999)).Return(nil, coreError.ErrNotFound)
			},
			wantErr:  true,
			wantCode: http.StatusNotFound,
		},
		{
			name: "delete error returns 500",
			ctx:  userCtx("user1"),
			id:   4,
			setup: func(r *repositories.MockVocabularyDeckRepo) {
				r.On("FindByID", mock.Anything, uint(4)).Return(userDeck, nil)
				r.On("Delete", mock.Anything, uint(4)).Return(errors.New("db error"))
			},
			wantErr:  true,
			wantCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(repositories.MockVocabularyDeckRepo)
			tt.setup(mockRepo)
			svc := NewVocabularyDeckService(mockRepo)

			err := svc.Delete(tt.ctx, tt.id)

			if tt.wantErr {
				assert.NotNil(t, err)
				assert.Equal(t, tt.wantCode, err.Code)
			} else {
				assert.Nil(t, err)
			}
			mockRepo.AssertExpectations(t)
		})
	}
}

func TestVocabularyDeckService_CreateAsSystem(t *testing.T) {
	tests := []struct {
		name      string
		body      req.CreateVocabularyDeckReq
		setup     func(*repositories.MockVocabularyDeckRepo)
		wantIsDefault bool
		wantErr   bool
		wantCode  int
	}{
		{
			name: "success — system deck has no user and is_default=true",
			body: req.CreateVocabularyDeckReq{Name: "System Deck", Level: "beginner"},
			setup: func(r *repositories.MockVocabularyDeckRepo) {
				r.On("Create", mock.Anything, mock.MatchedBy(func(d *models.VocabularyDeck) bool {
					return d.IsDefault && d.UserID == nil
				})).Run(func(args mock.Arguments) {
					args.Get(1).(*models.VocabularyDeck).ID = 10
				}).Return(nil)
			},
			wantIsDefault: true,
		},
		{
			name: "db error returns 500",
			body: req.CreateVocabularyDeckReq{Name: "Fail"},
			setup: func(r *repositories.MockVocabularyDeckRepo) {
				r.On("Create", mock.Anything, mock.Anything).Return(errors.New("db error"))
			},
			wantErr:  true,
			wantCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(repositories.MockVocabularyDeckRepo)
			tt.setup(mockRepo)
			svc := NewVocabularyDeckService(mockRepo)

			result, err := svc.CreateAsSystem(adminCtx(), tt.body)

			if tt.wantErr {
				assert.NotNil(t, err)
				assert.Equal(t, tt.wantCode, err.Code)
			} else {
				assert.Nil(t, err)
				assert.True(t, result.IsDefault)
				assert.Nil(t, result.UserID)
			}
			mockRepo.AssertExpectations(t)
		})
	}
}

func TestVocabularyDeckService_DeleteAsSystem(t *testing.T) {
	systemDeck := &models.VocabularyDeck{ID: 1, Name: "System", IsDefault: true}

	tests := []struct {
		name     string
		id       uint
		setup    func(*repositories.MockVocabularyDeckRepo)
		wantErr  bool
		wantCode int
	}{
		{
			name: "success",
			id:   1,
			setup: func(r *repositories.MockVocabularyDeckRepo) {
				r.On("FindByID", mock.Anything, uint(1)).Return(systemDeck, nil)
				r.On("Delete", mock.Anything, uint(1)).Return(nil)
			},
		},
		{
			name: "not found returns 404",
			id:   999,
			setup: func(r *repositories.MockVocabularyDeckRepo) {
				r.On("FindByID", mock.Anything, uint(999)).Return(nil, coreError.ErrNotFound)
			},
			wantErr:  true,
			wantCode: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(repositories.MockVocabularyDeckRepo)
			tt.setup(mockRepo)
			svc := NewVocabularyDeckService(mockRepo)

			err := svc.DeleteAsSystem(adminCtx(), tt.id)

			if tt.wantErr {
				assert.NotNil(t, err)
				assert.Equal(t, tt.wantCode, err.Code)
			} else {
				assert.Nil(t, err)
			}
			mockRepo.AssertExpectations(t)
		})
	}
}
