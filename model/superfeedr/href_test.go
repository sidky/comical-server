package superfeedr

import (
	"encoding/json"
	"testing"
)

var href = &Href{
	Title: "Parker Solar Probe",
	Href:  "https://xkcd.com/2262/",
	Rel:   "alternate",
	Type:  "text/html",
}

var hrefIncomplete = &Href{
	Title: "",
	Href:  "https://xkcd.com/2262/",
	Rel:   "alternate",
}

const serialized =
	`{"title":"Parker Solar Probe","href":"https://xkcd.com/2262/","rel":"alternate","type":"text/html"}`

const serializedIncomplete =
	`{"title":"","href":"https://xkcd.com/2262/","rel":"alternate"}`

func TestHrefSerialization(t *testing.T) {
	b, err := json.Marshal(href)
	if err != nil {
		t.Fatalf("Unable to serialize: %v", err)
		return
	}

	jsonText := string(b)

	if jsonText != serialized {
		t.Fatalf("Serialized values didn't match. expected '%v', actual '%v'", serialized, jsonText)
	}
}

func TestHrefDeserialization(t *testing.T) {
	d := Href{}

	err := json.Unmarshal([]byte(serialized), &d)
	if err != nil {
		t.Fatalf("Unable to deserialize: %v", err)
		return
	}

	if *href != d {
		t.Fatalf("Deserialized values didn't match. Expected '%v', Actual '%v'", href, d)
	}
}

func TestHrefIncompleteDeserialization(t *testing.T) {
	d := Href{}
	err := json.Unmarshal([]byte(serializedIncomplete), &d)
	if err != nil {
		t.Fatalf("Unable to deserialize: %v", err)
		return
	}

	if *hrefIncomplete != d {
		t.Fatalf("Deserialized values didn't match. Expected '%v', actual '%v'", hrefIncomplete, d)
	}
}