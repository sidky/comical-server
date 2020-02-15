package comics

import "time"

type Entry struct {
	Origin string
	PermalinkUrl string
	Title string
	Published time.Time
	Description string
	Images []string
}
