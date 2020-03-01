package storage

import (
	"comical/model/comics"
	"comical/model/store"
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
)

func NewComicsEntry(ctx context.Context, item *comics.Entry) error {

	entryName := fmt.Sprintf("comics.%s", item.Origin)

	r, err := GetClient().Collection("comics").Doc(item.Origin).Collection("entries").Doc(generatePath(item.ID)).Set(ctx, store.ToStoreEntry(item))
	log.Printf("Result: %v", r)
	if err != nil {
		return err
	}
	log.Printf("Successfully updated %s", entryName)
	return nil
}

func generatePath(id string) string {
	hasher := md5.New()
	hasher.Write([]byte(id))
	return hex.EncodeToString(hasher.Sum(nil))
}