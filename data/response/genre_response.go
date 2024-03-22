package response


type CreateNewGenreResponse struct {
	// ID   int    `json:"id" `
	Name string `json:"name" `
}


type UpdateGenreResponse struct {
	ID   int    `json:"id" `
	Name string `json:"name" `
}

type DeleteGenreResponse struct {
	Message string `json:"message"`
}


type GenreResponse struct {
	ID   int    `json:"id" `
	Name string `json:"name" `
}