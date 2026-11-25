package link

type LinkCreateRequest struct {
	URL string `json:"url" validate:"required,url"`
}

type LinkUpdateRequest struct {
	URL  string `json:"url" validate:"required,url"`
	Hash string `json:"hash"`
}

// GetAllLinksResponse содержит страницу ссылок и общее число записей, доступных для пагинации.
type GetAllLinksResponse struct {
	Links []Link `json:"links"`
	Count int64  `json:"count"`
}
