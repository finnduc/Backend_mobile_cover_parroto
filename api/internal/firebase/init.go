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
	if cfg.CredentialsFile == "" {
		log.Println("WARNING: Firebase credentials not configured — protected routes will be unavailable")
		return nil, errors.New("Firebase not configured")
	}

	ctx := context.Background()

	var app *firebase.App
	var err error

	app, err = firebase.NewApp(ctx, nil, option.WithAuthCredentialsFile(option.ServiceAccount, cfg.CredentialsFile))
	if err != nil {
		return nil, err
	}

	client, err := app.Auth(ctx)
	if err != nil {
		return nil, err
	}

	return &firebaseClient{auth: client}, nil
}
