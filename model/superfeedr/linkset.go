package superfeedr

type Linkset struct {
	Alternate  []Href `json:"alternate"`
	Superfeedr []Href `json:"superfeedr"`
	Self       []Href `json:"self"`
}
