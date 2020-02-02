package superfeedr

import (
	"encoding/json"
	"reflect"
	"testing"
)

var feedItem = FeedItem{
	ID:            "https://xkcd.com/2262/",
	Published:     1580446800,
	Updated:       1580446800,
	Title:         "Parker Solar Probe",
	Summary:       "<img src=\"https://imgs.xkcd.com/comics/parker_solar_probe.png\" title=\"It will get within 9 or 10 Sun-diameters of the &quot;bottom&quot; (the Sun's surface) which seems pretty far when you put it that way, but from up here on Earth it's practically all the way down.\" alt=\"It will get within 9 or 10 Sun-diameters of the &quot;bottom&quot; (the Sun's surface) which seems pretty far when you put it that way, but from up here on Earth it's practically all the way down.\" />",
	PermalinkUrl:  "https://xkcd.com/2262/",
	StandardLinks: Linkset{
		Alternate: []Href {
			{
				Title: "Parker Solar Probe",
				Href: "https://xkcd.com/2262/",
				Rel: "alternate",
				Type: "text/html",
			},
		},
	},
	ItemSource:    Source{
		ID: "xkcd-com-2020-1-31-5",
		Title: "xkcd.com",
		Updated: 1580446800,
		Published: 1580446800,
		PermalinkUrl: "https://xkcd.com/",
		StandardLinks: Linkset{
			Alternate: []Href{
				{
					Title: "xkcd.com",
					Href:  "https://xkcd.com/",
					Rel:   "alternate",
					Type:  "text/html",
				},
			},
			Superfeedr: []Href{
				{
					Title: "",
					Href:  "https://xkcd.com/rss.xml",
					Rel:   "superfeedr",
					Type: "",
				},
			},
		},
	},
	Language:      "en",
}

const serializedFeedItem =
	`{"id":"https://xkcd.com/2262/","published":1580446800,"updated":1580446800,"title":"Parker Solar Probe","summary":"\u003cimg src=\"https://imgs.xkcd.com/comics/parker_solar_probe.png\" title=\"It will get within 9 or 10 Sun-diameters of the \u0026quot;bottom\u0026quot; (the Sun's surface) which seems pretty far when you put it that way, but from up here on Earth it's practically all the way down.\" alt=\"It will get within 9 or 10 Sun-diameters of the \u0026quot;bottom\u0026quot; (the Sun's surface) which seems pretty far when you put it that way, but from up here on Earth it's practically all the way down.\" /\u003e","permalinkUrl":"https://xkcd.com/2262/","standardLinks":{"alternate":[{"title":"Parker Solar Probe","href":"https://xkcd.com/2262/","rel":"alternate","type":"text/html"}],"superfeedr":null},"source":{"id":"xkcd-com-2020-1-31-5","title":"xkcd.com","updated":1580446800,"published":1580446800,"permalinkUrl":"https://xkcd.com/","standardLinks":{"alternate":[{"title":"xkcd.com","href":"https://xkcd.com/","rel":"alternate","type":"text/html"}],"superfeedr":[{"title":"","href":"https://xkcd.com/rss.xml","rel":"superfeedr","type":""}]}},"language":"en"}`

func TestFeedItemSerialization(t *testing.T) {
	b, e := json.Marshal(feedItem)
	if e != nil {
		t.Fatalf("Unable to serialize: %v", e)
		return
	}

	jsonFeedItem := string(b)

	if serializedFeedItem != jsonFeedItem {
		t.Fatalf("Expected: '%v', Actual: '%v'", serializedFeedItem, jsonFeedItem)
	}

}

func TestFeedItemDeserialization(t *testing.T) {
	f := FeedItem{}
	e := json.Unmarshal([]byte(serializedFeedItem), &f)

	if e != nil {
		t.Fatalf("Unable to deserialize: %v", e)
		return
	}

	if !reflect.DeepEqual(feedItem, f) {
		t.Fatalf("Expected: '%v', Actual: '%v'", feedItem, f)
	}
}