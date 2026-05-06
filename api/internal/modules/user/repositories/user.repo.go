package repositories

import (
	"context"

	"go-cover-parroto/internal/core/enums"
	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/database"
	"go-cover-parroto/internal/database/models"
	"go-cover-parroto/internal/firebase"

	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/iterator"
)

type IUserRepo interface {
	FindByID(ctx context.Context, id string) (*models.User, error)
	FindAll(ctx context.Context, query *database.Query) (*response.PaginatedResult[*models.User], error)
	// Update(ctx context.Context, user *models.User) error
	Delete(ctx context.Context, id string) error
}

type userRepo struct {
	fbClient firebase.IFirebaseAuth
}

func NewUserRepo(fbClient firebase.IFirebaseAuth) IUserRepo {
	return &userRepo{fbClient: fbClient}
}

func (r *userRepo) FindByID(ctx context.Context, id string) (*models.User, error) {
	var user models.User

	userRecord, err := r.fbClient.GetUserByID(ctx, id)

	if err != nil {
		return nil, err
	}

	user.ID = userRecord.UID
	user.Email = userRecord.Email

	user.UserRole = enums.UserRoleUser
	if userRecord.CustomClaims != nil {
		if roleStr, ok := userRecord.CustomClaims[string(enums.CustomClaimKeyUserRole)].(string); ok {
			user.UserRole = enums.UserRole(roleStr)
		}
	}

	return &user, nil

}

func (r *userRepo) FindAll(ctx context.Context, query *database.Query) (*response.PaginatedResult[*models.User], error) {
	// 1. Normalize pagination parameters
	page := query.Page
	if page < 1 {
		page = 1
	}
	limit := query.Limit
	if limit < 1 {
		limit = 10
	}

	iter := r.fbClient.Users(ctx, "") // empty page token = start
	pager := iterator.NewPager(iter, limit, "")

	for i := 1; i < page; i++ {
		var discard []*auth.ExportedUserRecord
		_, err := pager.NextPage(&discard)
		if err == iterator.Done {
			// Requested page beyond last user – return empty result
			return &response.PaginatedResult[*models.User]{
				Data: []*models.User{},
				Meta: response.NewMeta(page, limit, 0), // total = 0
			}, nil
		}
		if err != nil {
			return nil, err
		}
	}

	var users []*auth.ExportedUserRecord
	_, err := pager.NextPage(&users)
	if err != nil && err != iterator.Done {
		return nil, err
	}

	resultData := make([]*models.User, 0, len(users))
	for _, fbUser := range users {
		userRole := enums.UserRoleUser // default
		if fbUser.CustomClaims != nil {
			if roleVal, ok := fbUser.CustomClaims[string(enums.CustomClaimKeyUserRole)]; ok {
				if roleStr, ok := roleVal.(string); ok {
					userRole = enums.UserRole(roleStr)
				}
			}
		}

		resultData = append(resultData, &models.User{
			ID:       fbUser.UID,
			Email:    fbUser.Email,
			UserRole: userRole,
		})
	}

	// 7. Build paginated result
	// Note: Firebase does not offer a cheap way to get total user count.
	// If you need exact Total and TotalPages, see "Counting users" below.
	// For now we set total = 0 (meaning "unknown") which yields TotalPages = 0.
	meta := response.NewMeta(page, limit, 0)

	// Optional: if you have a cached total count, use it here.
	// total, err := r.getCachedUserCount(ctx)
	// if err == nil { meta = response.NewMeta(page, limit, total) }

	return &response.PaginatedResult[*models.User]{
		Data: resultData,
		Meta: meta,
	}, nil
}
func (r *userRepo) Delete(ctx context.Context, id string) error {
	return r.fbClient.DeleteUser(ctx, id)
}
