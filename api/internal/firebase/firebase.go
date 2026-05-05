package firebase

import (
	"context"

	firebaseauth "firebase.google.com/go/v4/auth"
)

type IFirebaseAuth interface {
	VerifyIDToken(ctx context.Context, idToken string) (*firebaseauth.Token, error)
	SetCustomUserClaims(ctx context.Context, uid string, claims map[string]interface{}) error
	GetUserByID(ctx context.Context, uid string) (*firebaseauth.UserRecord, error)
	DeleteUser(ctx context.Context, uid string) error
	UpdateUser(ctx context.Context, uid string, user *firebaseauth.UserToUpdate) (ur *firebaseauth.UserRecord, err error)
	Users(ctx context.Context, nextPageToken string) *firebaseauth.UserIterator
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

func (f *firebaseClient) GetUserByID(ctx context.Context, uid string) (*firebaseauth.UserRecord, error) {
	return f.auth.GetUser(ctx, uid)
}

func (f *firebaseClient) DeleteUser(ctx context.Context, uid string) error {
	return f.auth.DeleteUser(ctx, uid)
}

func (f *firebaseClient) UpdateUser(ctx context.Context, uid string, user *firebaseauth.UserToUpdate) (ur *firebaseauth.UserRecord, err error) {
	return f.auth.UpdateUser(ctx, uid, user)
}

func (f *firebaseClient) Users(ctx context.Context, nextPageToken string) *firebaseauth.UserIterator {
	return f.auth.Users(ctx, nextPageToken)
}
