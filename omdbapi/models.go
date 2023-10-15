package omdbapi

type MovieSearched struct {
	Title  string
	Year   string
	ImdbID string `json:"ImdbID"`
	Type   string
	Poster string
}

type Movie struct {
	Title,
	Year,
	Rated,
	Genre,
	Actors,
	Writer,
	Plot,
	Country,
	Poster,
	Awards,
	ImdbRating,
	ImdbID,
	ImdbVotes,
	BoxOffice string
}

type SearchResult struct {
	Search       []MovieSearched
	TotalResults int `json:",string"`
}
