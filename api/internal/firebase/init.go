package firebase

import (
	"context"
	"errors"
	"go-cover-parroto/internal/configs"
	"log"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

func Init(cfg configs.FirebaseConfig) (IFirebaseAuth, error) {
	if cfg.CredentialsFile == "" && cfg.ProjectID == "" {
		log.Println("WARNING: Firebase credentials not configured — protected routes will be unavailable")
		return nil, errors.New("Firebase not configured")
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
