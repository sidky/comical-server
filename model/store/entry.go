package store

import "comical/model/comics"

type StoreEntry struct {
	ID          string   `firestore:"id"`
	Permalink   string   `firestore:"permalink,omitempty"`
	Title       string   `firestore:"title,omitempty"`
	Published   int64    `firestore:"published,omitempty"`
	Description string   `firestore:"description,omitempty"`
	Images      []string `firestore:"images,omitempty"`
}

func ToStoreEntry(entry *comics.Entry) StoreEntry {
	return StoreEntry{
		ID:          entry.ID,
		Permalink:   entry.PermalinkUrl,
		Title:       entry.Title,
		Published:   entry.Published.UnixNano(),
		Description: entry.Description,
		Images:      entry.Images,
	}
}
