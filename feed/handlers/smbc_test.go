package handlers

import (
	"comical/model/superfeedr"
	"testing"
)

var smbcFeedItem = superfeedr.FeedItem{
	ID:           "https://www.smbc-comics.com/comic/socializing",
	Published:    1580446800,
	Updated:      1580446800,
	Title:        "Saturday Morning Breakfast Cereal - Socializing",
	Summary:      "<a href=\"https://www.smbc-comics.com/comic/socializing\"><img src=\"https://www.smbc-comics.com/comics/1581782138-20200215.png\" /><br /><br />Click here to go see the bonus panel!</a><p>Hovertext:<br/>If you send hatemail about this, it's proof you were poorly socialized.</p><br />Today's News:<br />",
	PermalinkUrl: "https://www.smbc-comics.com/comic/socializing",
	StandardLinks: superfeedr.Linkset{
		Alternate: []superfeedr.Href{
			{
				Title: "Saturday Morning Breakfast Cereal - Socializing",
				Href:  "https://www.smbc-comics.com/comic/socializing",
				Rel:   "alternate",
				Type:  "text/html",
			},
		},
	},
	Language: "en",
}

func TestSmbcConvert(t *testing.T) {
	h := SmbcHandler{}

	item, err := h.Convert(&smbcFeedItem)

	if err != nil {
		t.Fail()
	}

	if item == nil {
		t.Fatal("item should be parsed")
	}

	if len(item.Images) != 1 {
		t.Fatalf("Expected 1 image, actual: %d images", len(item.Images))
	}

	src := item.Images[0]
	if src != "https://www.smbc-comics.com/comics/1581782138-20200215.png" {
		t.Fatalf("Wrong image: %s", src)
	}

	desc := item.Description
	if desc != "If you send hatemail about this, it's proof you were poorly socialized." {
		t.Fatalf("Wrong description: %s", desc)
	}
}
