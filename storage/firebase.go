package storage

import (
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
	"log"
	"os"
	"sync"
)

var firestoreClient *firestore.Client
var mu sync.Mutex

func InitStorage() (*firestore.Client, error) {

	mu.Lock()
	defer mu.Unlock()

	if firestoreClient != nil {
		return firestoreClient, nil
	}

	credentialBytes := os.Getenv("FIREBASE_CRED")

	opt := option.WithCredentialsJSON([]byte(credentialBytes))
	ctx := context.Background()
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, err
	}
	firestoreClient, err = app.Firestore(ctx)
	if err != nil {
		return nil, err
	}

	log.Printf("Initialized Firestore client")

	return firestoreClient, nil
}

func GetClient() *firestore.Client {
	mu.Lock()
	defer mu.Unlock()

	return firestoreClient
}
