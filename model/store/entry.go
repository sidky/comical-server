package store

import "comical/model/comics"

type StoreEntry struct {
	Permalink string `firestore:"permalink,omitempty"`
	Title string `firestore:"title,omitempty"`
	Published int64 `firestore:"published,omitempty"`
	Description string `firestore:"description,omitempty"`
	Images []string `firestore:"images,omitempty"`
}

func ToStoreEntry(entry *comics.Entry) StoreEntry {
	return StoreEntry{
		Permalink:   entry.PermalinkUrl,
		Title:       entry.Title,
		Published:   entry.Published.Unix(),
		Description: entry.Description,
		Images:      entry.Images,
	}
}
