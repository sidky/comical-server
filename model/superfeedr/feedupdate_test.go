package superfeedr

import (
	"encoding/json"
	"reflect"
	"testing"
)

var feedUpdate = FeedUpdate{
	UpdateStatus: Status{
		Code:              304,
		Http:              "Fetched (ring) 306 900",
		NextFetch:         1580581304,
		Velocity:          0.5,
		Title:             "xkcd.com",
		Period:            900,
		LastFetch:         1580580404,
		LastParse:         1580565988,
		LastMaintenanceAt: 1580565988,
		Feed:              "ttps://xkcd.com/rss.xml",
	},
	Title:   "xkcd.com",
	Updated: 0,
	ID:      "",
	Items: []FeedItem{
		{
			ID:           "https://xkcd.com/2262/",
			Published:    1580446800,
			Updated:      1580446800,
			Title:        "Parker Solar Probe",
			Summary:      "<img src=\"https://imgs.xkcd.com/comics/parker_solar_probe.png\" title=\"It will get within 9 or 10 Sun-diameters of the &quot;bottom&quot; (the Sun's surface) which seems pretty far when you put it that way, but from up here on Earth it's practically all the way down.\" alt=\"It will get within 9 or 10 Sun-diameters of the &quot;bottom&quot; (the Sun's surface) which seems pretty far when you put it that way, but from up here on Earth it's practically all the way down.\" />",
			PermalinkUrl: "https://xkcd.com/2262/",
			StandardLinks: Linkset{
				Alternate: []Href{
					{
						Title: "Parker Solar Probe",
						Href:  "https://xkcd.com/2262/",
						Rel:   "alternate",
						Type:  "text/html",
					},
				},
				Superfeedr: nil,
			},
			ItemSource: Source{
				ID:           "xkcd-com-2020-1-31-5",
				Title:        "xkcd.com",
				Updated:      1580446800,
				Published:    1580446800,
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
						},
					},
				},
			},
			Language: "en",
		},
	},
}

const realJson = `{
    "status": {
        "code": 304,
        "http": "Fetched (ring) 306 900",
        "nextFetch": 1580581304,
        "velocity": 0.5,
        "title": "xkcd.com",
        "period": 900,
        "lastFetch": 1580580404,
        "lastParse": 1580565988,
        "lastMaintenanceAt": 1580565988,
        "feed": "ttps://xkcd.com/rss.xml"
    },
    "title": "xkcd.com",
    "updated": null,
    "id": "",
    "items": [
        {
            "id": "https://xkcd.com/2262/",
            "published": 1580446800,
            "updated": 1580446800,
            "title": "Parker Solar Probe",
            "summary": "<img src=\"https://imgs.xkcd.com/comics/parker_solar_probe.png\" title=\"It will get within 9 or 10 Sun-diameters of the &quot;bottom&quot; (the Sun's surface) which seems pretty far when you put it that way, but from up here on Earth it's practically all the way down.\" alt=\"It will get within 9 or 10 Sun-diameters of the &quot;bottom&quot; (the Sun's surface) which seems pretty far when you put it that way, but from up here on Earth it's practically all the way down.\" />",
            "permalinkUrl": "https://xkcd.com/2262/",
            "standardLinks": {
                "alternate": [
                    {
                        "title": "Parker Solar Probe",
                        "href": "https://xkcd.com/2262/",
                        "rel": "alternate",
                        "type": "text/html"
                    }
                ]
            },
            "source": {
                "id": "xkcd-com-2020-1-31-5",
                "title": "xkcd.com",
                "updated": 1580446800,
                "published": 1580446800,
                "permalinkUrl": "https://xkcd.com/",
                "standardLinks": {
                    "alternate": [
                        {
                            "title": "xkcd.com",
                            "href": "https://xkcd.com/",
                            "rel": "alternate",
                            "type": "text/html"
                        }
                    ],
                    "superfeedr": [
                        {
                            "title": "",
                            "href": "https://xkcd.com/rss.xml",
                            "rel": "superfeedr"
                        }
                    ]
                }
            },
            "language": "en"
        }
    ]
}`

func TestFeedUpdateSerialization(t *testing.T) {
	u := FeedUpdate{}
	err := json.Unmarshal([]byte(realJson), &u)
	if err != nil {
		t.Fatalf("Unable to parse: %v", err)
		return
	}
	if !reflect.DeepEqual(feedUpdate, u) {
		t.Fatalf("Expected: %v, actual: %v", feedUpdate, u)
	}
}
