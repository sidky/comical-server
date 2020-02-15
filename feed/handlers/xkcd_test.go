package handlers

import (
	"comical/model/superfeedr"
	"testing"
)

var feedItem = superfeedr.FeedItem{
	ID:           "https://xkcd.com/2262/",
	Published:    1580446800,
	Updated:      1580446800,
	Title:        "Parker Solar Probe",
	Summary:      "<img src=\"https://imgs.xkcd.com/comics/parker_solar_probe.png\" title=\"It will get within 9 or 10 Sun-diameters of the &quot;bottom&quot; (the Sun's surface) which seems pretty far when you put it that way, but from up here on Earth it's practically all the way down.\" alt=\"It will get within 9 or 10 Sun-diameters of the &quot;bottom&quot; (the Sun's surface) which seems pretty far when you put it that way, but from up here on Earth it's practically all the way down.\" />",
	PermalinkUrl: "https://xkcd.com/2262/",
	StandardLinks: superfeedr.Linkset{
		Alternate: []superfeedr.Href{
			{
				Title: "Parker Solar Probe",
				Href:  "https://xkcd.com/2262/",
				Rel:   "alternate",
				Type:  "text/html",
			},
		},
	},
	ItemSource: superfeedr.Source{
		ID:           "xkcd-com-2020-1-31-5",
		Title:        "xkcd.com",
		Updated:      1580446800,
		Published:    1580446800,
		PermalinkUrl: "https://xkcd.com/",
		StandardLinks: superfeedr.Linkset{
			Alternate: []superfeedr.Href{
				{
					Title: "xkcd.com",
					Href:  "https://xkcd.com/",
					Rel:   "alternate",
					Type:  "text/html",
				},
			},
			Superfeedr: []superfeedr.Href{
				{
					Title: "",
					Href:  "https://xkcd.com/rss.xml",
					Rel:   "superfeedr",
					Type:  "",
				},
			},
		},
	},
	Language: "en",
}

func TestXkcdConverter(t *testing.T) {
	h := XkcdFeedHandler{}

	item, err := h.Convert(feedItem)

	if err != nil {
		t.Fail()
	}

	if item == nil {
		t.Fatal("item should be parsed, and not be nil")
	}

	if len(item.Images) != 1 {
		t.Fatalf("Expected 1 image, actual: %d images", len(item.Images))
	}

	src := item.Images[0]
	if src != "https://imgs.xkcd.com/comics/parker_solar_probe.png" {
		t.Fatalf("Parsed image url didn't match")
	}

}
