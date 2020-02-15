package superfeedr

import (
	"encoding/json"
	"testing"
)

var status = &Status{
	Code:              304,
	Http:              "Fetched (ring) 306 900",
	NextFetch:         1580581304,
	Velocity:          0.5,
	Title:             "xkcd.com",
	Period:            900,
	LastFetch:         1580580404,
	LastParse:         1580565988,
	LastMaintenanceAt: 33881,
	Feed:              "https://xkcd.com/rss.xm",
}

const expectedSerialized = `{"code":304,"http":"Fetched (ring) 306 900","nextFetch":1580581304,"velocity":0.5,"title":"xkcd.com","period":900,"lastFetch":1580580404,"lastParse":1580565988,"lastMaintenanceAt":33881,"feed":"https://xkcd.com/rss.xm"}`

func TestStatusSerialization(t *testing.T) {
	b, err := json.Marshal(status)
	if err != nil {
		t.Fatalf("Unable to serialize: %v", err)
	}
	serializedJson := string(b)

	if serializedJson != expectedSerialized {
		t.Fatalf("Serialized didn't match. expected '%s', actual='%s'", expectedSerialized, serializedJson)
	}
}

func TestStatusDeserialization(t *testing.T) {
	var d = Status{}
	err := json.Unmarshal([]byte(expectedSerialized), &d)
	if err != nil {
		t.Fatalf("Unable to deserialize: %v", err)
	}

	if d != *status {
		t.Fatalf("Deserialized values didn't match. expected '%v', actual '%v'", status, d)
	}
}
