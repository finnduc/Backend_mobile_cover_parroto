package firebase

import (
	"context"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"go-cover-parroto/internal/configs"
	"google.golang.org/api/option"
)

var AuthClient *auth.Client

func Init(cfg configs.FirebaseConfig) error {
	var app *firebase.App
	var err error

	ctx := context.Background()

	if cfg.CredentialsFile != "" {
		app, err = firebase.NewApp(ctx, nil, option.WithCredentialsFile(cfg.CredentialsFile))
	} else {
		conf := &firebase.Config{ProjectID: cfg.ProjectID}
		app, err = firebase.NewApp(ctx, conf)
	}

	if err != nil {
		return err
	}

	AuthClient, err = app.Auth(ctx)
	return err
}
