package request

type CreateNewDirectorRequest struct {
	Name    string `json:"name" form:"name" validate:"required"`
	MovieID int    `json:"movie_id" form:"movie_id" validate:"required"`
	AwardID int    `json:"award_id" form:"award_id" validate:"required"`
}

type UpdateDirectorRequest struct {
	ID      int    `validate:"required"`
	Name    string `json:"name" form:"name" validate:"required"`
	MovieID int    `json:"movie_id" form:"movie_id" validate:"required"`
	AwardID int    `json:"award_id" form:"award_id" validate:"required"`
}
