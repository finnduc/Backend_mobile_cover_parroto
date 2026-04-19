package firebase

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	firebase "firebase.google.com/go/v4"
	firebaseauth "firebase.google.com/go/v4/auth"
	"go-cover-parroto/internal/configs"
	"google.golang.org/api/option"
)

type TokenResult struct {
	IDToken      string
	RefreshToken string
	ExpiresIn    string
	Email        string
}

type IFirebaseAuth interface {
	VerifyIDToken(ctx context.Context, idToken string) (*firebaseauth.Token, error)
	SignIn(ctx context.Context, email, password string) (*TokenResult, error)
}

// --- real client ---

type firebaseClient struct {
	auth      *firebaseauth.Client
	webAPIKey string
}

func (f *firebaseClient) VerifyIDToken(ctx context.Context, idToken string) (*firebaseauth.Token, error) {
	return f.auth.VerifyIDToken(ctx, idToken)
}

func (f *firebaseClient) SignIn(ctx context.Context, email, password string) (*TokenResult, error) {
	url := fmt.Sprintf("https://identitytoolkit.googleapis.com/v1/accounts:signInWithPassword?key=%s", f.webAPIKey)
	payload := fmt.Sprintf(`{"email":%q,"password":%q,"returnSecureToken":true}`, email, password)

	req, _ := http.NewRequestWithContext(ctx, http.MethodPost, url, strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var raw struct {
		IDToken      string `json:"idToken"`
		RefreshToken string `json:"refreshToken"`
		ExpiresIn    string `json:"expiresIn"`
		Email        string `json:"email"`
		Error        *struct {
			Message string `json:"message"`
		} `json:"error"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&raw); err != nil {
		return nil, err
	}
	if raw.Error != nil {
		return nil, errors.New(raw.Error.Message)
	}

	return &TokenResult{
		IDToken:      raw.IDToken,
		RefreshToken: raw.RefreshToken,
		ExpiresIn:    raw.ExpiresIn,
		Email:        raw.Email,
	}, nil
}

// --- stub (no credentials configured) ---

type stubFirebaseAuth struct{}

func (s *stubFirebaseAuth) VerifyIDToken(_ context.Context, _ string) (*firebaseauth.Token, error) {
	return nil, errors.New("firebase not configured")
}

func (s *stubFirebaseAuth) SignIn(_ context.Context, _, _ string) (*TokenResult, error) {
	return nil, errors.New("firebase not configured")
}

// --- Init ---

func EmptyConfig() configs.FirebaseConfig {
	return configs.FirebaseConfig{}
}

func Init(cfg configs.FirebaseConfig) (IFirebaseAuth, error) {
	if cfg.CredentialsFile == "" && cfg.ProjectID == "" {
		log.Println("WARNING: Firebase credentials not configured — protected routes will be unavailable")
		return &stubFirebaseAuth{}, nil
	}

	ctx := context.Background()

	var app *firebase.App
	var err error

	if cfg.CredentialsFile != "" {
		app, err = firebase.NewApp(ctx, nil, option.WithCredentialsFile(cfg.CredentialsFile))
	} else {
		conf := &firebase.Config{ProjectID: cfg.ProjectID}
		app, err = firebase.NewApp(ctx, conf)
	}
	if err != nil {
		return nil, err
	}

	client, err := app.Auth(ctx)
	if err != nil {
		return nil, err
	}

	return &firebaseClient{auth: client, webAPIKey: cfg.WebAPIKey}, nil
}
