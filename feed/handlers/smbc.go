package handlers

import (
	"comical/model/comics"
	"comical/model/superfeedr"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
	"strings"
	"time"
)

type SmbcHandler struct {
}

func NewSmbcHandler() FeedItemHandler {
	return SmbcHandler{}
}

func (s SmbcHandler) Name() string {
	panic("smbc-comics")
}

func (s SmbcHandler) CanHandle(item *superfeedr.FeedItem) bool {
	return strings.Index(item.ID, "https://www.smbc-comics.com/") == 0
}

func (s SmbcHandler) Convert(item *superfeedr.FeedItem) (*comics.Entry, error) {
	nodes, err := html.ParseFragment(strings.NewReader(item.Summary), nil)
	if err != nil {
		return nil, err
	}

	if len(nodes) > 1 {
		return nil, fmt.Errorf("smbc: %d nodes in HTML, not sure how to parse: '%s'", len(nodes), item.Summary)
	}

	var entry *comics.Entry = nil

	for _, node := range nodes {
		document := goquery.NewDocumentFromNode(node)
		image := document.Find("img").First()
		description := ""

		if image != nil {
			src, _ := image.Attr("src")
			document.Find("p").EachWithBreak(func(i int, s *goquery.Selection) bool {
				if strings.Contains(s.Text(), "Hovertext:") {
					index := strings.Index(s.Text(), "Hovertext:")
					description = s.Text()[index+len("Hovertext:"):]
					return false
				} else {
					return true
				}
			})

			entry = &comics.Entry{
				ID:           item.ID,
				Origin:       "smbc",
				PermalinkUrl: item.PermalinkUrl,
				Title:        item.Title,
				Published:    time.Unix(item.Published/1000, (item.Published%1000)*1000000),
				Description:  description,
				Images:       []string{src},
			}
		}
	}
	return entry, nil
}

var _ FeedItemHandler = (*SmbcHandler)(nil)
