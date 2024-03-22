package request

type CreateNewMoviesRequest struct {
	Title   string `json:"title" form:"title" validate:"required"`
	Year    int    `json:"year" form:"year" validate:"required"`
	GenreID int    `json:"genre_id" form:"genre_id" validate:"required"`
	AwardID int    `json:"award_id" form:"award_id" validate:"required"`
}

type UpdateMoviesRequest struct {
	ID      int    `validate:"required"`
	Title   string `json:"title" form:"title" validate:"required"`
	Year    int    `json:"year" form:"year" validate:"required"`
	GenreID int    `json:"genre_id" form:"genre_id" validate:"required"`
	AwardID int    `json:"award_id" form:"award_id" validate:"required"`
}
