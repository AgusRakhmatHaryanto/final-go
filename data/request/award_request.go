package request

type CreateNewAwardRequest struct {
	Title string `json:"title" form:"title" validate:"required"`
	Year  int    `json:"year" form:"year" validate:"required"`
}

type UpdateAwardRequest struct {
	ID    int    `validate:"required"`
	Title string `json:"title" form:"title" validate:"required"`
	Year  int    `json:"year" form:"year" validate:"required"`
}
