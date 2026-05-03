package firebase

import (
	"context"

	firebaseauth "firebase.google.com/go/v4/auth"
)

type TokenResult struct {
	IDToken      string
	RefreshToken string
	ExpiresIn    string
	Email        string
}

type IFirebaseAuth interface {
	VerifyIDToken(ctx context.Context, idToken string) (*firebaseauth.Token, error)
}
type firebaseClient struct {
	auth      *firebaseauth.Client
	webAPIKey string
}

func (f *firebaseClient) VerifyIDToken(ctx context.Context, idToken string) (*firebaseauth.Token, error) {
	return f.auth.VerifyIDToken(ctx, idToken)
}
