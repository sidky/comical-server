package superfeedr

type Href struct {
	Title string `json:"title"`
	Href  string `json:"href"`
	Rel   string `json:"rel"`
	Type  string `json:"type"`
}
