package comics

import "time"

type Entry struct {
	ID           string
	Origin       string
	PermalinkUrl string
	Title        string
	Published    time.Time
	Description  string
	Images       []string
}
