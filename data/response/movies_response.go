package response

type CreateMovieResponse struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Year    int    `json:"year"`
	AwardID int    `json:"award_id"`
	GenreID int    `json:"genre_id"`
}

type UpdateMovieResponse struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Year    int    `json:"year"`
	AwardID int    `json:"award_id"`
	GenreID int    `json:"genre_id"`
}

type MovieGetAwardResponse struct {
	ID   int    `json:"id"`
	Title string `json:"title"`
	Year int    `json:"year"`
}

type MovieGetGenreResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type MovieGetAllResponse struct {
	ID      int                   `json:"id"`
	Title   string                `json:"title"`
	Year    int                   `json:"year"`
	AwardID int                   `json:"award_id"`
	GenreID int                   `json:"genre_id"`
	Award   AwardResponse
	Genre   GenreResponse
}

type DeleteMovieResponse struct {
	Message string `json:"message"`
}
