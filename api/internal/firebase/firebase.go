package firebase

import (
	"context"
	"errors"
	"log"

	firebase "firebase.google.com/go/v4"
	firebaseauth "firebase.google.com/go/v4/auth"
	"go-cover-parroto/internal/configs"
	"google.golang.org/api/option"
)

type IFirebaseAuth interface {
	VerifyIDToken(ctx context.Context, idToken string) (*firebaseauth.Token, error)
}

// stubFirebaseAuth is used when no Firebase credentials are configured.
// All token verifications will fail with unauthorized.
type stubFirebaseAuth struct{}

func (s *stubFirebaseAuth) VerifyIDToken(_ context.Context, _ string) (*firebaseauth.Token, error) {
	return nil, errors.New("firebase not configured")
}

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

	return app.Auth(ctx)
}
