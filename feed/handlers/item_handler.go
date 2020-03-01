package handlers

import (
	"comical/model/comics"
	"comical/model/superfeedr"
	"fmt"
)

// Add all comics handlers here
var handlers = [...]FeedItemHandler{
	NewXkcdHandler(), // XKCD
	NewSmbcHandler(), // Saturday Morning Breakfast Cereal
}

func ExtractEntry(item *superfeedr.FeedItem) (*comics.Entry, error) {
	for _, handler := range handlers {
		if handler.CanHandle(item) {
			entry, err := handler.Convert(item)

			if err != nil {
				return nil, err
			} else {
				return entry, nil
			}
		}
	}

	return nil, fmt.Errorf("unable to find a handler for %v", item)
}
