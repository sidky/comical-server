package superfeedr

import (
	"encoding/json"
	"reflect"
	"testing"
)

var source = Source{
	ID:            "xkcd-com-2020-1-31-5",
	Title:         "xkcd.com",
	Updated:       1580446800,
	Published:     1580446800,
	PermalinkUrl:  "https://xkcd.com/",
	StandardLinks: Linkset{
		Alternate: []Href{
			{
				Title: "xkcd.com",
				Href: "https://xkcd.com",
				Rel: "alternate",
				Type: "text/html",
			},
		},
		Superfeedr: []Href {
			{
				Title: "",
				Href: "https://xkcd.com/rss",
				Rel: "superfeedr",
			},
		},
	},
}

const serializedSource =
	`{"id":"xkcd-com-2020-1-31-5","title":"xkcd.com","updated":1580446800,"published":1580446800,"permalinkUrl":"https://xkcd.com/","standardLinks":{"alternate":[{"title":"xkcd.com","href":"https://xkcd.com","rel":"alternate","type":"text/html"}],"superfeedr":[{"title":"","href":"https://xkcd.com/rss","rel":"superfeedr","type":""}]}}`

func TestSourceSerialization(t *testing.T) {
	b, err := json.Marshal(&source)
	if err != nil {
		t.Fatalf("Unable to serialize: %v", err)
		return
	}

	jsonSource := string(b)
	if serializedSource != jsonSource {
		t.Fatalf("Didn't match. Expected: '%v', Actual: '%v'", serializedSource, jsonSource)
	}
}

func TestSourceDeserialization(t *testing.T) {
	s := Source{}

	err := json.Unmarshal([]byte(serializedSource), &s)
	if err != nil {
		t.Fatalf("Unable to deserialize: %v", err)
		return
	}

	if !reflect.DeepEqual(source, s) {
		t.Fatalf("Didn't match. Expected: '%v', Actual: '%v'", source, s)
	}
}