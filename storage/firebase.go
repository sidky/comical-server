package storage

import (
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go"
)

func InitStorage() (*firestore.Client, error) {
	ctx := context.Background()
	conf := &firebase.Config{ProjectID:"comical-266916"}
	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		return nil, err
	}
	client, err := app.Firestore(ctx)
	if err != nil {
		return nil, err
	}
	return client, nil
}
