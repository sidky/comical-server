package storage

import (
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
	"os"
)

func InitStorage() (*firestore.Client, error) {

	credentialBytes := os.Getenv("FIREBASE_CRED")

	opt := option.WithCredentialsJSON([]byte(credentialBytes))
	ctx := context.Background()
	//conf := &firebase.Config{ProjectID:"comical-266916"}
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, err
	}
	client, err := app.Firestore(ctx)
	if err != nil {
		return nil, err
	}
	return client, nil
}
