package models

type Movie struct {
	ID         int      `json:"id"`
	Title      string   `json:"title"`
	Poster     string   `json:"poster"`
	Year       string   `json:"year"`
	Rated      string   `json:"rated"`
	Released   string   `json:"released"`
	Runtime    string   `json:"runtime"`
	Director   string   `json:"director"`
	Writer     string   `json:"writer"`
	Actors     string   `json:"actors"`
	Plot       string   `json:"plot"`
	Country    string   `json:"country"`
	Awards     string   `json:"awards"`
	Metascore  string   `json:"metascore"`
	IMDBRating string   `json:"imdb_rating"`
	IMDBVotes  string   `json:"imdb_votes"`
	IMDBID     string   `json:"imdb_id"`
	Type       string   `json:"type"`
	Genres     []string `json:"genres"`
	Images     []string `json:"images"`
}