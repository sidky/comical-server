package superfeedr

type FeedUpdate struct {
	UpdateStatus Status     `json:"status"`
	Title        string     `json:"title"`
	Updated      int64      `json:"updated"`
	ID           string     `json:"id"`
	Items        []FeedItem `json:"items"`
}
