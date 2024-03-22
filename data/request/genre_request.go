package request

type CreateNewGenreRequest struct {
	Name string `json:"name" form:"name" validate:"required"`
}

type UpdateGenreRequest struct {
	ID   int    `validate:"required"`
	Name string `json:"name" form:"name" validate:"required"`
}