package superfeedr

type Source struct {
	ID            string  `json:"id"`
	Title         string  `json:"title"`
	Updated       int64   `json:"updated"`
	Published     int64   `json:"published"`
	PermalinkUrl  string  `json:"permalinkUrl"`
	StandardLinks Linkset `json:"standardLinks"`
}
