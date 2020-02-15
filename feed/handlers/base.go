package handlers

import (
	"comical/model/comics"
	"comical/model/superfeedr"
)

type FeedItemHandler interface {
	Name() string
	CanHandle(item *superfeedr.FeedItem) bool
	Convert(item *superfeedr.FeedItem) (*comics.Entry, error)
}
