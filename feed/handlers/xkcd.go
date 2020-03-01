package handlers

import (
	"comical/model/comics"
	"comical/model/superfeedr"
	"fmt"
	"golang.org/x/net/html"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type XkcdFeedHandler struct {
}

func NewXkcdHandler() XkcdFeedHandler {
	return XkcdFeedHandler{}
}

func (h XkcdFeedHandler) Name() string {
	return "xkcd"
}

func (h XkcdFeedHandler) CanHandle(item *superfeedr.FeedItem) bool {
	return strings.Index(item.ID, "https://xkcd.com/") == 0
}

func (h XkcdFeedHandler) Convert(item *superfeedr.FeedItem) (*comics.Entry, error) {
	nodes, err := html.ParseFragment(strings.NewReader(item.Summary), nil)
	if err != nil {
		return nil, err
	}

	if len(nodes) > 1 {
		return nil, fmt.Errorf("xkcd: %d nodes in HTML, not sure how to parse: '%s'", len(nodes), item.Summary)
	}

	var entry *comics.Entry;
	entry = nil

	for _, node := range nodes {
		document := goquery.NewDocumentFromNode(node)
		image := document.Find("img").First()

		if image != nil {
			src, _ := image.Attr("src")
			alt, _ := image.Attr("alt")

			entry = &comics.Entry{
				ID:           item.ID,
				Origin:       h.Name(),
				PermalinkUrl: item.PermalinkUrl,
				Title:        item.Title,
				Published:    time.Unix(item.Published/1000, (item.Published%1000)*1000000),
				Description:  alt,
				Images: []string{
					src,
				},
			}
		}
	}
	return entry, nil
}

// Ensure XkcdFeedHandler implements FeedItemHandler
var _ FeedItemHandler = (*XkcdFeedHandler)(nil)
