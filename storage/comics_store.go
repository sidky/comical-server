package storage

import (
	"comical/model/comics"
	"comical/model/store"
	"context"
	"fmt"
	"log"
)

func NewComicsEntry(ctx context.Context, item *comics.Entry) error {

	entryName := fmt.Sprintf("comics.%s", item.Origin)

	log.Printf("client: %v\n", firestoreClient)
	log.Printf("comics/: %v\n", firestoreClient.Collection("comics"))
	log.Printf("comics/%s/: %v\n", item.Origin, firestoreClient.Collection("comics").Doc(item.Origin))

	_, _, err := firestoreClient.Collection("comics").Doc(item.Origin).Collection("entries").Add(ctx, store.ToStoreEntry(item))
	if err != nil {
		return err
	}
	log.Printf("Successfully updated %s", entryName)
	return nil
}
