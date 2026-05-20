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
	"go-cover-parroto/internal/modules/vocabulary_item/dtos/req"
	itemrepo "go-cover-parroto/internal/modules/vocabulary_item/repositories"
	deckrepo "go-cover-parroto/internal/modules/vocabulary_deck/repositories"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func userCtx(userID string) context.Context {
	ctx := context.WithValue(context.Background(), enums.ContextKeyUserID, userID)
	return context.WithValue(ctx, enums.ContextKeyUserRole, enums.UserRoleUser)
}

func strPtr(s string) *string { return &s }

func TestVocabularyItemService_List(t *testing.T) {
	deckID := uint(1)
	tests := []struct {
		name     string
		query    req.ListVocabularyItemQuery
		setup    func(*itemrepo.MockVocabularyItemRepo, *deckrepo.MockVocabularyDeckRepo)
		wantLen  int
		wantErr  bool
		wantCode int
	}{
		{
			name:  "success — returns items for deck",
			query: req.ListVocabularyItemQuery{DeckID: &deckID, Page: 1, Limit: 10},
			setup: func(ir *itemrepo.MockVocabularyItemRepo, dr *deckrepo.MockVocabularyDeckRepo) {
				ir.On("FindAll", mock.Anything, mock.Anything).Return(
					&response.PaginatedResult[*models.VocabularyItem]{
						Data: []*models.VocabularyItem{
							{ID: 1, DeckID: 1, Phrase: "give up", Meaning: "Từ bỏ"},
							{ID: 2, DeckID: 1, Phrase: "look forward to", Meaning: "Mong chờ"},
						},
						Meta: response.NewMeta(1, 10, 2),
					}, nil)
			},
			wantLen: 2,
		},
		{
			name:  "success — empty deck returns empty list",
			query: req.ListVocabularyItemQuery{DeckID: &deckID, Page: 1, Limit: 10},
			setup: func(ir *itemrepo.MockVocabularyItemRepo, dr *deckrepo.MockVocabularyDeckRepo) {
				ir.On("FindAll", mock.Anything, mock.Anything).Return(
					&response.PaginatedResult[*models.VocabularyItem]{
						Data: []*models.VocabularyItem{},
						Meta: response.NewMeta(1, 10, 0),
					}, nil)
			},
			wantLen: 0,
		},
		{
			name:  "db error returns 500",
			query: req.ListVocabularyItemQuery{DeckID: &deckID, Page: 1, Limit: 10},
			setup: func(ir *itemrepo.MockVocabularyItemRepo, dr *deckrepo.MockVocabularyDeckRepo) {
				ir.On("FindAll", mock.Anything, mock.Anything).Return(nil, errors.New("db error"))
			},
			wantErr:  true,
			wantCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockItemRepo := new(itemrepo.MockVocabularyItemRepo)
			mockDeckRepo := new(deckrepo.MockVocabularyDeckRepo)
			tt.setup(mockItemRepo, mockDeckRepo)
			svc := NewVocabularyItemService(mockItemRepo, mockDeckRepo)

			result, err := svc.List(context.Background(), tt.query)

			if tt.wantErr {
				assert.NotNil(t, err)
				assert.Equal(t, tt.wantCode, err.Code)
			} else {
				assert.Nil(t, err)
				assert.Len(t, result.Data, tt.wantLen)
			}
			mockItemRepo.AssertExpectations(t)
		})
	}
}

func TestVocabularyItemService_GetByID(t *testing.T) {
	tests := []struct {
		name     string
		id       uint
		setup    func(*itemrepo.MockVocabularyItemRepo, *deckrepo.MockVocabularyDeckRepo)
		wantID   uint
		wantErr  bool
		wantCode int
	}{
		{
			name: "success",
			id:   1,
			setup: func(ir *itemrepo.MockVocabularyItemRepo, dr *deckrepo.MockVocabularyDeckRepo) {
				ir.On("FindByID", mock.Anything, uint(1)).Return(
					&models.VocabularyItem{ID: 1, DeckID: 1, Phrase: "give up"}, nil)
			},
			wantID: 1,
		},
		{
			name: "not found returns 404",
			id:   999,
			setup: func(ir *itemrepo.MockVocabularyItemRepo, dr *deckrepo.MockVocabularyDeckRepo) {
				ir.On("FindByID", mock.Anything, uint(999)).Return(nil, coreError.ErrNotFound)
			},
			wantErr:  true,
			wantCode: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockItemRepo := new(itemrepo.MockVocabularyItemRepo)
			mockDeckRepo := new(deckrepo.MockVocabularyDeckRepo)
			tt.setup(mockItemRepo, mockDeckRepo)
			svc := NewVocabularyItemService(mockItemRepo, mockDeckRepo)

			result, err := svc.GetByID(context.Background(), tt.id)

			if tt.wantErr {
				assert.NotNil(t, err)
				assert.Equal(t, tt.wantCode, err.Code)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.wantID, result.ID)
			}
			mockItemRepo.AssertExpectations(t)
		})
	}
}

func TestVocabularyItemService_Create(t *testing.T) {
	ownerID := "user1"
	userDeck := &models.VocabularyDeck{ID: 1, Name: "My Deck", UserID: &ownerID}
	systemDeck := &models.VocabularyDeck{ID: 2, Name: "System Deck", IsDefault: true}

	tests := []struct {
		name     string
		ctx      context.Context
		body     req.CreateVocabularyItemFromDeckReq
		setup    func(*itemrepo.MockVocabularyItemRepo, *deckrepo.MockVocabularyDeckRepo)
		wantErr  bool
		wantCode int
	}{
		{
			name: "success — owner adds item to own deck",
			ctx:  userCtx("user1"),
			body: req.CreateVocabularyItemFromDeckReq{DeckID: 1, Phrase: "give up", NormalizedPhrase: "give up", Meaning: "Từ bỏ"},
			setup: func(ir *itemrepo.MockVocabularyItemRepo, dr *deckrepo.MockVocabularyDeckRepo) {
				dr.On("FindByID", mock.Anything, uint(1)).Return(userDeck, nil)
				ir.On("Create", mock.Anything, mock.MatchedBy(func(i *models.VocabularyItem) bool {
					return i.Phrase == "give up" && i.DeckID == 1
				})).Run(func(args mock.Arguments) {
					args.Get(1).(*models.VocabularyItem).ID = 9
				}).Return(nil)
			},
		},
		{
			name: "success — user adds item to system deck",
			ctx:  userCtx("user1"),
			body: req.CreateVocabularyItemFromDeckReq{DeckID: 2, Phrase: "run out", NormalizedPhrase: "run out", Meaning: "Hết"},
			setup: func(ir *itemrepo.MockVocabularyItemRepo, dr *deckrepo.MockVocabularyDeckRepo) {
				dr.On("FindByID", mock.Anything, uint(2)).Return(systemDeck, nil)
				ir.On("Create", mock.Anything, mock.Anything).Return(nil)
			},
		},
		{
			name:     "unauthenticated returns 401",
			ctx:      context.Background(),
			body:     req.CreateVocabularyItemFromDeckReq{DeckID: 1, Phrase: "test"},
			setup:    func(ir *itemrepo.MockVocabularyItemRepo, dr *deckrepo.MockVocabularyDeckRepo) {},
			wantErr:  true,
			wantCode: http.StatusUnauthorized,
		},
		{
			name: "deck not found returns 404",
			ctx:  userCtx("user1"),
			body: req.CreateVocabularyItemFromDeckReq{DeckID: 999, Phrase: "test"},
			setup: func(ir *itemrepo.MockVocabularyItemRepo, dr *deckrepo.MockVocabularyDeckRepo) {
				dr.On("FindByID", mock.Anything, uint(999)).Return(nil, coreError.ErrNotFound)
			},
			wantErr:  true,
			wantCode: http.StatusNotFound,
		},
		{
			name: "forbidden — non-owner adds to user deck returns 403",
			ctx:  userCtx("other_user"),
			body: req.CreateVocabularyItemFromDeckReq{DeckID: 1, Phrase: "test"},
			setup: func(ir *itemrepo.MockVocabularyItemRepo, dr *deckrepo.MockVocabularyDeckRepo) {
				dr.On("FindByID", mock.Anything, uint(1)).Return(userDeck, nil)
			},
			wantErr:  true,
			wantCode: http.StatusForbidden,
		},
		{
			name: "db error on create returns 500",
			ctx:  userCtx("user1"),
			body: req.CreateVocabularyItemFromDeckReq{DeckID: 1, Phrase: "test", NormalizedPhrase: "test", Meaning: "test"},
			setup: func(ir *itemrepo.MockVocabularyItemRepo, dr *deckrepo.MockVocabularyDeckRepo) {
				dr.On("FindByID", mock.Anything, uint(1)).Return(userDeck, nil)
				ir.On("Create", mock.Anything, mock.Anything).Return(errors.New("db error"))
			},
			wantErr:  true,
			wantCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockItemRepo := new(itemrepo.MockVocabularyItemRepo)
			mockDeckRepo := new(deckrepo.MockVocabularyDeckRepo)
			tt.setup(mockItemRepo, mockDeckRepo)
			svc := NewVocabularyItemService(mockItemRepo, mockDeckRepo)

			result, err := svc.Create(tt.ctx, tt.body)

			if tt.wantErr {
				assert.NotNil(t, err)
				assert.Equal(t, tt.wantCode, err.Code)
				assert.Nil(t, result)
			} else {
				assert.Nil(t, err)
				assert.NotNil(t, result)
			}
			mockItemRepo.AssertExpectations(t)
			mockDeckRepo.AssertExpectations(t)
		})
	}
}

func TestVocabularyItemService_Update(t *testing.T) {
	ownerID := "user1"
	itemWithUserDeck := &models.VocabularyItem{
		ID:     1,
		DeckID: 1,
		Phrase: "old phrase",
		Deck:   &models.VocabularyDeck{ID: 1, UserID: &ownerID},
	}
	itemWithSystemDeck := &models.VocabularyItem{
		ID:     2,
		DeckID: 2,
		Phrase: "system item",
		Deck:   &models.VocabularyDeck{ID: 2, IsDefault: true},
	}

	tests := []struct {
		name       string
		ctx        context.Context
		id         uint
		body       req.UpdateVocabularyItemReq
		setup      func(*itemrepo.MockVocabularyItemRepo, *deckrepo.MockVocabularyDeckRepo)
		wantPhrase string
		wantErr    bool
		wantCode   int
	}{
		{
			name: "success — owner updates own deck item",
			ctx:  userCtx("user1"),
			id:   1,
			body: req.UpdateVocabularyItemReq{Phrase: "new phrase", NormalizedPhrase: "new phrase", Meaning: "meaning"},
			setup: func(ir *itemrepo.MockVocabularyItemRepo, dr *deckrepo.MockVocabularyDeckRepo) {
				ir.On("FindByID", mock.Anything, uint(1)).Return(itemWithUserDeck, nil)
				ir.On("Update", mock.Anything, mock.MatchedBy(func(i *models.VocabularyItem) bool {
					return i.Phrase == "new phrase"
				})).Return(nil)
			},
			wantPhrase: "new phrase",
		},
		{
			name: "success — any user updates system deck item",
			ctx:  userCtx("any_user"),
			id:   2,
			body: req.UpdateVocabularyItemReq{Phrase: "updated", NormalizedPhrase: "updated", Meaning: "meaning"},
			setup: func(ir *itemrepo.MockVocabularyItemRepo, dr *deckrepo.MockVocabularyDeckRepo) {
				ir.On("FindByID", mock.Anything, uint(2)).Return(itemWithSystemDeck, nil)
				ir.On("Update", mock.Anything, mock.Anything).Return(nil)
			},
			wantPhrase: "updated",
		},
		{
			name: "forbidden — non-owner updates user deck item",
			ctx:  userCtx("other_user"),
			id:   1,
			body: req.UpdateVocabularyItemReq{Phrase: "hack"},
			setup: func(ir *itemrepo.MockVocabularyItemRepo, dr *deckrepo.MockVocabularyDeckRepo) {
				ir.On("FindByID", mock.Anything, uint(1)).Return(itemWithUserDeck, nil)
			},
			wantErr:  true,
			wantCode: http.StatusForbidden,
		},
		{
			name: "not found returns 404",
			ctx:  userCtx("user1"),
			id:   999,
			body: req.UpdateVocabularyItemReq{Phrase: "x"},
			setup: func(ir *itemrepo.MockVocabularyItemRepo, dr *deckrepo.MockVocabularyDeckRepo) {
				ir.On("FindByID", mock.Anything, uint(999)).Return(nil, coreError.ErrNotFound)
			},
			wantErr:  true,
			wantCode: http.StatusNotFound,
		},
		{
			name: "db update error returns 500",
			ctx:  userCtx("user1"),
			id:   1,
			body: req.UpdateVocabularyItemReq{Phrase: "x", NormalizedPhrase: "x", Meaning: "x"},
			setup: func(ir *itemrepo.MockVocabularyItemRepo, dr *deckrepo.MockVocabularyDeckRepo) {
				ir.On("FindByID", mock.Anything, uint(1)).Return(itemWithUserDeck, nil)
				ir.On("Update", mock.Anything, mock.Anything).Return(errors.New("db error"))
			},
			wantErr:  true,
			wantCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockItemRepo := new(itemrepo.MockVocabularyItemRepo)
			mockDeckRepo := new(deckrepo.MockVocabularyDeckRepo)
			tt.setup(mockItemRepo, mockDeckRepo)
			svc := NewVocabularyItemService(mockItemRepo, mockDeckRepo)

			result, err := svc.Update(tt.ctx, tt.id, tt.body)

			if tt.wantErr {
				assert.NotNil(t, err)
				assert.Equal(t, tt.wantCode, err.Code)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.wantPhrase, result.Phrase)
			}
			mockItemRepo.AssertExpectations(t)
		})
	}
}

func TestVocabularyItemService_Delete(t *testing.T) {
	ownerID := "user1"
	itemWithUserDeck := &models.VocabularyItem{
		ID:   1,
		Deck: &models.VocabularyDeck{ID: 1, UserID: &ownerID},
	}
	itemWithSystemDeck := &models.VocabularyItem{
		ID:   2,
		Deck: &models.VocabularyDeck{ID: 2, IsDefault: true},
	}

	tests := []struct {
		name     string
		ctx      context.Context
		id       uint
		setup    func(*itemrepo.MockVocabularyItemRepo, *deckrepo.MockVocabularyDeckRepo)
		wantErr  bool
		wantCode int
	}{
		{
			name: "success — owner deletes own deck item",
			ctx:  userCtx("user1"),
			id:   1,
			setup: func(ir *itemrepo.MockVocabularyItemRepo, dr *deckrepo.MockVocabularyDeckRepo) {
				ir.On("FindByID", mock.Anything, uint(1)).Return(itemWithUserDeck, nil)
				ir.On("Delete", mock.Anything, uint(1)).Return(nil)
			},
		},
		{
			name: "success — deletes system deck item",
			ctx:  userCtx("any_user"),
			id:   2,
			setup: func(ir *itemrepo.MockVocabularyItemRepo, dr *deckrepo.MockVocabularyDeckRepo) {
				ir.On("FindByID", mock.Anything, uint(2)).Return(itemWithSystemDeck, nil)
				ir.On("Delete", mock.Anything, uint(2)).Return(nil)
			},
		},
		{
			name: "forbidden — non-owner returns 403",
			ctx:  userCtx("other_user"),
			id:   1,
			setup: func(ir *itemrepo.MockVocabularyItemRepo, dr *deckrepo.MockVocabularyDeckRepo) {
				ir.On("FindByID", mock.Anything, uint(1)).Return(itemWithUserDeck, nil)
			},
			wantErr:  true,
			wantCode: http.StatusForbidden,
		},
		{
			name: "not found returns 404",
			ctx:  userCtx("user1"),
			id:   999,
			setup: func(ir *itemrepo.MockVocabularyItemRepo, dr *deckrepo.MockVocabularyDeckRepo) {
				ir.On("FindByID", mock.Anything, uint(999)).Return(nil, coreError.ErrNotFound)
			},
			wantErr:  true,
			wantCode: http.StatusNotFound,
		},
		{
			name: "delete error returns 500",
			ctx:  userCtx("user1"),
			id:   1,
			setup: func(ir *itemrepo.MockVocabularyItemRepo, dr *deckrepo.MockVocabularyDeckRepo) {
				ir.On("FindByID", mock.Anything, uint(1)).Return(itemWithUserDeck, nil)
				ir.On("Delete", mock.Anything, uint(1)).Return(errors.New("db error"))
			},
			wantErr:  true,
			wantCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockItemRepo := new(itemrepo.MockVocabularyItemRepo)
			mockDeckRepo := new(deckrepo.MockVocabularyDeckRepo)
			tt.setup(mockItemRepo, mockDeckRepo)
			svc := NewVocabularyItemService(mockItemRepo, mockDeckRepo)

			err := svc.Delete(tt.ctx, tt.id)

			if tt.wantErr {
				assert.NotNil(t, err)
				assert.Equal(t, tt.wantCode, err.Code)
			} else {
				assert.Nil(t, err)
			}
			mockItemRepo.AssertExpectations(t)
		})
	}
}
