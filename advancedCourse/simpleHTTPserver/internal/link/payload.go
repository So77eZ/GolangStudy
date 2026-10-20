package link

type LinkCreateRequest struct {
	URL string `json:"url" validate:"required,url"`
}

type LinkCreateResponse struct {
	URL string `json:"url" validate:"required,url"`
}
