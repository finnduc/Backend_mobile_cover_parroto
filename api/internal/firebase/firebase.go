package firebase

import (
	"context"

	firebaseauth "firebase.google.com/go/v4/auth"
)

type IFirebaseAuth interface {
	VerifyIDToken(ctx context.Context, idToken string) (*firebaseauth.Token, error)
	SetCustomUserClaims(ctx context.Context, uid string, claims map[string]interface{}) error
}
type firebaseClient struct {
	auth *firebaseauth.Client
}

func (f *firebaseClient) VerifyIDToken(ctx context.Context, idToken string) (*firebaseauth.Token, error) {
	return f.auth.VerifyIDToken(ctx, idToken)
}

func (f *firebaseClient) SetCustomUserClaims(ctx context.Context, uid string, claims map[string]interface{}) error {
	return f.auth.SetCustomUserClaims(ctx, uid, claims)
}
