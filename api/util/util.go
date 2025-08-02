package util

import (
	"context"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

var FirebaseApp *firebase.App

func InitFirebase(ctx context.Context) error {
	opt := option.WithCredentialsFile("path/to/your/firebase-credentials.json") //TODO: replace this with actual firebase credentials
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return err
	}
	FirebaseApp = app
	return nil
}
