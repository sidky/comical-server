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

	_, _, err := GetClient().Collection("comics").Doc(item.Origin).Collection("entries").Add(ctx, store.ToStoreEntry(item))
	if err != nil {
		return err
	}
	log.Printf("Successfully updated %s", entryName)
	return nil
}
