package album

type Album struct {
	Title    string  `json:"title"`
	Artist   string  `json:"artist"`
	Price    float64 `json:"price"`
	ImageURL string  `json:"imageUrl"`
}

type AlbumResponse struct {
	ID       string  `json:"id"`
	Title    string  `json:"title"`
	Artist   string  `json:"artist"`
	Price    float64 `json:"price"`
	ImageURL string  `json:"imageUrl"`
}
