package superfeedr

import (
	"encoding/json"
	"reflect"
	"testing"
)

var linkset = &Linkset{
	Alternate: []Href{
		{
			Title: "xkcd.com",
			Href:  "https://xkcd.com",
			Rel:   "alternate",
			Type:  "text/html",
		},
		{
			Title: "xkcd.com",
			Href:  "https://m.xkcd.com",
			Rel:   "alternate",
			Type:  "text/html",
		},
	},
	Superfeedr: []Href{
		{
			Title: "",
			Href:  "https://xkcd.com/rss.xml",
		},
	},
}

const serializedLinkset = `{"alternate":[{"title":"xkcd.com","href":"https://xkcd.com","rel":"alternate","type":"text/html"},{"title":"xkcd.com","href":"https://m.xkcd.com","rel":"alternate","type":"text/html"}],"superfeedr":[{"title":"","href":"https://xkcd.com/rss.xml","rel":"","type":""}]}`

func TestLinksetSerialization(t *testing.T) {
	b, err := json.Marshal(linkset)
	if err != nil {
		t.Fatalf("Unable to serialize: %v", err)
		return
	}

	json := string(b)

	if serializedLinkset != json {
		t.Fatalf("Serialized linkset didn't match. Expected: '%v', Actual: '%v'",
			serializedLinkset, json)
	}
}

func TestLinksetDeserialization(t *testing.T) {
	l := Linkset{}
	err := json.Unmarshal([]byte(serializedLinkset), &l)
	if err != nil {
		t.Fatalf("Unable to deserialize: %v", err)
		return
	}

	if reflect.DeepEqual(linkset, l) {
		t.Fatalf("Deserialized linkset didn't match. Expected: '%v', actual: '%v'",
			linkset, l)
	}
}
