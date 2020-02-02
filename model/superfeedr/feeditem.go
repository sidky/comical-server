package superfeedr

type FeedItem struct {
	ID string `json:"id"`
	Published int64 `json:"published"`
	Updated int64 `json:"updated"`
	Title string `json:"title"`
	Summary string `json:"summary"`
	PermalinkUrl string `json:"permalinkUrl"`
	StandardLinks Linkset `json:"standardLinks"`
	ItemSource Source `json:"source"`
	Language string `json:"language"`
}
