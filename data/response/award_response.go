package response

type CreateNewAwardResponse struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Year  int    `json:"year"`
}

type UpdateAwardResponse struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Year  int    `json:"year"`
}

type DeleteAwardResponse struct {
	Message string `json:"message"`
}

type AwardResponse struct {
	ID    int    `json:"id"`
	Title string `json:"name"`
	Year  int    `json:"year"`
}
