package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	coreError "go-cover-parroto/internal/core/errors"
	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/database/models"
	"go-cover-parroto/internal/firebase"
	authreq "go-cover-parroto/internal/modules/auth/dtos/req"
	authres "go-cover-parroto/internal/modules/auth/dtos/res"
	"go-cover-parroto/internal/modules/auth/repositories"
)

type IAuthService interface {
	SyncUser(ctx context.Context, firebaseToken string) (*authres.SyncRes, *response.AppError)
	GetToken(ctx context.Context, apiKey string, body authreq.GetTokenReq) (*authres.TokenRes, *response.AppError)
}

type authService struct {
	repo   repositories.IAuthRepo
	fbAuth firebase.IFirebaseAuth
}

func NewAuthService(repo repositories.IAuthRepo, fbAuth firebase.IFirebaseAuth) IAuthService {
	return &authService{repo: repo, fbAuth: fbAuth}
}

func (s *authService) SyncUser(ctx context.Context, firebaseToken string) (*authres.SyncRes, *response.AppError) {
	decoded, err := s.fbAuth.VerifyIDToken(ctx, firebaseToken)
	if err != nil {
		return nil, response.Unauthorized("invalid firebase token")
	}

	email, _ := decoded.Claims["email"].(string)
	name, _ := decoded.Claims["name"].(string)
	picture, _ := decoded.Claims["picture"].(string)

	if email == "" {
		return nil, response.BadRequest("firebase token missing email")
	}

	user, err := s.repo.FindByEmail(ctx, email)
	if err != nil {
		if !errors.Is(err, coreError.ErrNotFound) {
			return nil, response.Internal("failed to find user")
		}

		user = &models.User{
			Email:     email,
			Name:      name,
			AvatarURL: picture,
			Password:  "",
		}

		if createErr := s.repo.Create(ctx, user); createErr != nil {
			return nil, response.Internal("failed to create user")
		}
	}

	return &authres.SyncRes{
		ID:        user.ID,
		Email:     user.Email,
		Name:      user.Name,
		AvatarURL: user.AvatarURL,
	}, nil
}

func (s *authService) GetToken(ctx context.Context, apiKey string, body authreq.GetTokenReq) (*authres.TokenRes, *response.AppError) {
	url := fmt.Sprintf("https://identitytoolkit.googleapis.com/v1/accounts:signInWithPassword?key=%s", apiKey)
	payload := fmt.Sprintf(`{"email":%q,"password":%q,"returnSecureToken":true}`, body.Email, body.Password)

	req, _ := http.NewRequestWithContext(ctx, http.MethodPost, url, strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, response.Internal("failed to contact Firebase")
	}
	defer resp.Body.Close()

	var result struct {
		IDToken      string `json:"idToken"`
		RefreshToken string `json:"refreshToken"`
		ExpiresIn    string `json:"expiresIn"`
		Email        string `json:"email"`
		Error        *struct {
			Message string `json:"message"`
		} `json:"error"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, response.Internal("failed to parse Firebase response")
	}
	if result.Error != nil {
		return nil, response.Unauthorized(result.Error.Message)
	}

	return &authres.TokenRes{
		IDToken:      result.IDToken,
		RefreshToken: result.RefreshToken,
		ExpiresIn:    result.ExpiresIn,
		Email:        result.Email,
	}, nil
}
