package response

type CreateNewDirectorResponse struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	MovieID int    `json:"movie_id"`
}

type UpdateDirectorResponse struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	MovieID int    `json:"movie_id"`
}

type DirectorGetMovieResponse struct {
	ID      int                   `json:"id"`
	Title   string                `json:"title"`
	Year    int                   `json:"year"`
	AwardID int                   `json:"award_id"`
	GenreID int                   `json:"genre_id"`
	Award   MovieGetAwardResponse `json:"award"`
	Genre   MovieGetGenreResponse `json:"genre"`
}

type DirectorGetAllResponse struct {
	ID      int                      `json:"id"`
	Name    string                   `json:"name"`
	MovieID int                      `json:"movie_id"`
	Movie   MovieGetAllResponse
}

type DeleteDirectorResponse struct {
	Message string `json:"message"`
}