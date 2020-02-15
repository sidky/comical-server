package superfeedr

type Status struct {
	Code              int16   `json:"code"`
	Http              string  `json:"http"`
	NextFetch         int64   `json:"nextFetch"`
	Velocity          float32 `json:"velocity"`
	Title             string  `json:"title"`
	Period            int32   `json:"period"`
	LastFetch         int64   `json:"lastFetch"`
	LastParse         int64   `json:"lastParse"`
	LastMaintenanceAt int64   `json:"lastMaintenanceAt"`
	Feed              string  `json:"feed"`
}
